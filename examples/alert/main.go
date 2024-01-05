package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/core/proto/alert"
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

	listResp, err := client.ListAlert(ctx, &alert.ListAlertRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListAlert API.", slog.Int("len", len(listResp.Alert)))

	listCondResp, err := client.ListAlertCondition(ctx, &alert.ListAlertConditionRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListAlertCondition API.", slog.Int("len", len(listCondResp.AlertCondition)))

	conditionID := listCondResp.AlertCondition[0].AlertConditionId
	putResp, err := client.PutAlert(ctx, &alert.PutAlertRequest{
		ProjectId: projectID,
		Alert: &alert.AlertForUpsert{
			AlertConditionId: conditionID,
			ProjectId:        projectID,
			Description:      "test",
			Severity:         "low", // "high", "medium", "low"
			Status:           alert.Status_ACTIVE,
		},
	})
	handleError(err)
	logger.Info("Success PutAlert API.", slog.Int("alert_id", int(putResp.Alert.AlertId)))

	err = client.AnalyzeAlert(ctx, &alert.AnalyzeAlertRequest{
		ProjectId:        projectID,
		AlertConditionId: []uint32{conditionID},
	})
	handleError(err)
	logger.Info("Success AnalyzeAlert API.")
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
