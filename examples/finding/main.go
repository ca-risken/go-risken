package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/core/proto/finding"
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
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	projectID := resp.ProjectID

	listResp, err := client.ListFinding(ctx, &finding.ListFindingRequest{
		ProjectId: projectID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success ListFinding API.", slog.Uint64("total", uint64(listResp.Total)))

	putResp, err := client.PutFinding(ctx, &finding.PutFindingRequest{
		ProjectId: projectID,
		Finding: &finding.FindingForUpsert{
			Description:      "desc",
			DataSource:       "ds",
			DataSourceId:     "dsid",
			ResourceName:     "rn",
			ProjectId:        projectID,
			OriginalMaxScore: 1.0,
			OriginalScore:    1.0,
			Data:             `{"test":true}`,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	findingID := putResp.Finding.FindingId
	logger.Info("Success PutFinding API.", slog.Uint64("id", findingID))

	getResp, err := client.GetFinding(ctx, &finding.GetFindingRequest{
		ProjectId: projectID,
		FindingId: findingID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success GetFinding API.", slog.Uint64("id", getResp.Finding.FindingId))

	err = client.DeleteFinding(ctx, &finding.DeleteFindingRequest{
		ProjectId: projectID,
		FindingId: findingID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success DeleteFinding API.", slog.Uint64("id", findingID))
}
