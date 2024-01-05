// AUTOGENERATED FILE BY CODE GENERATOR. DO NOT EDIT.
// Source: generator.yaml
package risken

import (
	"context"

	proto "github.com/ca-risken/datasource-api/proto/diagnosis"
)

func (c *Client) ListWpscanSetting(ctx context.Context, req *proto.ListWpscanSettingRequest) (*proto.ListWpscanSettingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/diagnosis/list-wpscan-setting", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListWpscanSettingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListPortscanSetting(ctx context.Context, req *proto.ListPortscanSettingRequest) (*proto.ListPortscanSettingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/diagnosis/list-portscan-setting", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListPortscanSettingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListPortscanTarget(ctx context.Context, req *proto.ListPortscanTargetRequest) (*proto.ListPortscanTargetResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/diagnosis/list-portscan-target", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListPortscanTargetResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListApplicationScan(ctx context.Context, req *proto.ListApplicationScanRequest) (*proto.ListApplicationScanResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/diagnosis/list-application-scan", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.ListApplicationScanResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetApplicationScanBasicSetting(ctx context.Context, req *proto.GetApplicationScanBasicSettingRequest) (*proto.GetApplicationScanBasicSettingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/diagnosis/get-application-scan-basic-setting", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.GetApplicationScanBasicSettingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetDiagnosisDataSource(ctx context.Context, req *proto.GetDiagnosisDataSourceRequest) (*proto.GetDiagnosisDataSourceResponse, error) {
	httpReq, err := c.NewRequest(ctx, "GET", "/api/v1/diagnosis/get-datasource", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.GetDiagnosisDataSourceResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) InvokeScanDiagnosis(ctx context.Context, req *proto.InvokeScanRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/invoke-scan", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) PutWpscanSetting(ctx context.Context, req *proto.PutWpscanSettingRequest) (*proto.PutWpscanSettingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/put-wpscan-setting", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutWpscanSettingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteWpscanSetting(ctx context.Context, req *proto.DeleteWpscanSettingRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/delete-wpscan-setting", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) PutApplicationScan(ctx context.Context, req *proto.PutApplicationScanRequest) (*proto.PutApplicationScanResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/put-application-scan", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutApplicationScanResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteApplicationScan(ctx context.Context, req *proto.DeleteApplicationScanRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/delete-application-scan", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) PutApplicationScanBasicSetting(ctx context.Context, req *proto.PutApplicationScanBasicSettingRequest) (*proto.PutApplicationScanBasicSettingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/put-application-scan-basic-setting", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutApplicationScanBasicSettingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteApplicationScanBasicSetting(ctx context.Context, req *proto.DeleteApplicationScanBasicSettingRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/delete-application-scan-basic-setting", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) PutPortscanSetting(ctx context.Context, req *proto.PutPortscanSettingRequest) (*proto.PutPortscanSettingResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/put-portscan-setting", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutPortscanSettingResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PutPortscanTarget(ctx context.Context, req *proto.PutPortscanTargetRequest) (*proto.PutPortscanTargetResponse, error) {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/put-portscan-target", req)
	if err != nil {
		return nil, err
	}
	httpResp, err := c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var resp proto.PutPortscanTargetResponse
	if err := decodeBodyWithDataKey(httpResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeletePortscanSetting(ctx context.Context, req *proto.DeletePortscanSettingRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/delete-portscan-setting", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeletePortscanTarget(ctx context.Context, req *proto.DeletePortscanTargetRequest) error {
	httpReq, err := c.NewRequest(ctx, "POST", "/api/v1/diagnosis/delete-portscan-target", req)
	if err != nil {
		return err
	}
	if _, err := c.Do(httpReq); err != nil {
		return err
	}
	return nil
}
