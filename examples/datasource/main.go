package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/datasource-api/proto/datasource"
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

	attackFlowResp, err := client.AnalyzeAttackFlow(ctx, &datasource.AnalyzeAttackFlowRequest{
		ProjectId:    projectID,
		CloudType:    "aws",
		CloudId:      "123456789012",
		ResourceName: "arn:aws:s3:::bucket-name",
	})
	handleError(err)
	logger.Info("Success AnalyzeAttackFlow API.", slog.Any("resp", attackFlowResp))
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
