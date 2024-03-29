// AUTOGENERATED FILE BY CODE GENERATOR. DO NOT EDIT.
// Source: generator.yaml
package risken

import (
	"context"

	proto "github.com/ca-risken/datasource-api/proto/datasource"
)

func (c *Client) AnalyzeAttackFlow(ctx context.Context, req *proto.AnalyzeAttackFlowRequest) (*proto.AnalyzeAttackFlowResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/datasource/get-attack-flow-analysis", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.AnalyzeAttackFlowResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
