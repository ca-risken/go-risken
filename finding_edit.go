package risken

import (
	"context"

	proto "github.com/ca-risken/core/proto/finding"
)

func (c *Client) ListFinding(ctx context.Context, req *proto.ListFindingRequest) (*proto.ListFindingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/finding/list-finding", req)
	if err != nil {
		return nil, err
	}

	// クエリを上書き
	query := httpReq.URL.Query()
	switch req.Status {
	case proto.FindingStatus_FINDING_UNKNOWN:
		query.Set("status", "0") // all
	case proto.FindingStatus_FINDING_ACTIVE:
		query.Set("status", "1") // active
	case proto.FindingStatus_FINDING_PENDING:
		query.Set("status", "2") // pending
	default:
		query.Set("status", "1") // デフォルトはactive
	}
	httpReq.URL.RawQuery = query.Encode()

	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListFindingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
