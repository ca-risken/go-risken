package risken

import (
	"context"

	proto "github.com/ca-risken/core/proto/alert"
)

func (c *Client) ListAlert(ctx context.Context, req *proto.ListAlertRequest) (*proto.ListAlertResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-alert", req)
	if err != nil {
		return nil, err
	}
	// クエリを上書き
	query := httpReq.URL.Query()
	if req.Status != nil {
		for _, status := range req.Status {
			switch status {
			case proto.Status_ACTIVE:
				query.Set("status", "1") // active
			case proto.Status_PENDING:
				query.Set("status", "2") // pending
			case proto.Status_DEACTIVE:
				query.Set("status", "3") // deactive
			}
		}
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListAlertResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
