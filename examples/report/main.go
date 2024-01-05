package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/ca-risken/core/proto/report"
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

	now := time.Now()
	fromDate := now.AddDate(0, 0, -1).Format("2006-01-02")
	toDate := now.Format("2006-01-02")
	reportResp, err := client.GetReportFinding(ctx, &report.GetReportFindingRequest{
		ProjectId: projectID,
		FromDate:  fromDate,
		ToDate:    toDate,
	})
	handleError(err)
	logger.Info("Success GetReportFinding API.", slog.Int("len", len(reportResp.ReportFinding)))
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
