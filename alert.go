// AUTOGENERATED FILE BY CODE GENERATOR. DO NOT EDIT.
// Source: generator.yaml
package risken

import (
	"context"

	proto "github.com/ca-risken/core/proto/alert"
)


func (c *Client) ListAlertHistory(ctx context.Context, req *proto.ListAlertHistoryRequest) (*proto.ListAlertHistoryResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-history", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListAlertHistoryResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListRelAlertFinding(ctx context.Context, req *proto.ListRelAlertFindingRequest) (*proto.ListRelAlertFindingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-rel_alert_finding", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListRelAlertFindingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAlertCondition(ctx context.Context, req *proto.ListAlertConditionRequest) (*proto.ListAlertConditionResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-condition", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListAlertConditionResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAlertRule(ctx context.Context, req *proto.ListAlertRuleRequest) (*proto.ListAlertRuleResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-rule", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListAlertRuleResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAlertCondRule(ctx context.Context, req *proto.ListAlertCondRuleRequest) (*proto.ListAlertCondRuleResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-condition_rule", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListAlertCondRuleResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListNotification(ctx context.Context, req *proto.ListNotificationRequest) (*proto.ListNotificationResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-notification", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListNotificationResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAlertCondNotification(ctx context.Context, req *proto.ListAlertCondNotificationRequest) (*proto.ListAlertCondNotificationResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/alert/list-condition_notification", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListAlertCondNotificationResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutAlert(ctx context.Context, req *proto.PutAlertRequest) (*proto.PutAlertResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-alert", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutAlertResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutAlertFirstViewedAt(ctx context.Context, req *proto.PutAlertFirstViewedAtRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-alert-first-viewed-at", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) PutAlertCondition(ctx context.Context, req *proto.PutAlertConditionRequest) (*proto.PutAlertConditionResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-condition", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutAlertConditionResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutAlertRule(ctx context.Context, req *proto.PutAlertRuleRequest) (*proto.PutAlertRuleResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-rule", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutAlertRuleResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutAlertCondRule(ctx context.Context, req *proto.PutAlertCondRuleRequest) (*proto.PutAlertCondRuleResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-condition_rule", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutAlertCondRuleResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutNotification(ctx context.Context, req *proto.PutNotificationRequest) (*proto.PutNotificationResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-notification", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutNotificationResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutAlertCondNotification(ctx context.Context, req *proto.PutAlertCondNotificationRequest) (*proto.PutAlertCondNotificationResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/put-condition_notification", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutAlertCondNotificationResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteAlertCondition(ctx context.Context, req *proto.DeleteAlertConditionRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/delete-condition", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteAlertRule(ctx context.Context, req *proto.DeleteAlertRuleRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/delete-rule", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteAlertCondRule(ctx context.Context, req *proto.DeleteAlertCondRuleRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/delete-condition_rule", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteNotification(ctx context.Context, req *proto.DeleteNotificationRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/delete-notification", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteAlertCondNotification(ctx context.Context, req *proto.DeleteAlertCondNotificationRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/delete-condition_notification", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) AnalyzeAlert(ctx context.Context, req *proto.AnalyzeAlertRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/analyze-alert", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) TestNotification(ctx context.Context, req *proto.TestNotificationRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/alert/test-notification", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}
