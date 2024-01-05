package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/datasource-api/proto/diagnosis"
	"github.com/ca-risken/go-risken"
)

const (
	RISKEN_ENDPOINT = "http://localhost:8001"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	ctx := context.Background()
	token := os.Getenv("RISKEN_API_TOKEN")
	client := risken.NewClient(token, risken.WithAPIEndpoint(RISKEN_ENDPOINT))

	resp, err := client.Signin(ctx)
	handleError(err)
	projectID := resp.ProjectID
	logger.Info("Success Signin.", slog.Int("project_id", int(projectID)))

	listResp, err := client.ListWpscanSetting(ctx, &diagnosis.ListWpscanSettingRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListWpscanSetting API.", slog.Int("len", len(listResp.WpscanSetting)))

	putResp, err := client.PutWpscanSetting(ctx, &diagnosis.PutWpscanSettingRequest{
		ProjectId: projectID,
		WpscanSetting: &diagnosis.WpscanSettingForUpsert{
			ProjectId:             projectID,
			DiagnosisDataSourceId: 1,
			TargetUrl:             "https://example.com",
			Status:                diagnosis.Status_CONFIGURED,
			Options:               `{"test":true}`,
		},
	})
	handleError(err)
	logger.Info("Success PutWpscanSetting API.", slog.Any("wpscan", putResp))

	err = client.DeleteWpscanSetting(ctx, &diagnosis.DeleteWpscanSettingRequest{
		ProjectId:       projectID,
		WpscanSettingId: putResp.WpscanSetting.WpscanSettingId,
	})
	handleError(err)
	logger.Info("Success DeleteWpscanSetting API.")
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
