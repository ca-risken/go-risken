package risken

import (
	"context"
	"net/http"
)

type SigninResponse struct {
	ProjectID     uint32 `json:"project_id,omitempty"`
	AccessTokenID uint32 `json:"access_token_id,omitempty"`
}

func (c *Client) Signin(ctx context.Context) (*SigninResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, "/api/v1/signin", nil)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var resp SigninResponse
	if err := decodeBody(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
