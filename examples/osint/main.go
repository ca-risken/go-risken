package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/datasource-api/proto/osint"
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

	listResp, err := client.ListOsint(ctx, &osint.ListOsintRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListOSINT API.", slog.Int("len", len(listResp.Osint)))

	putResp, err := client.PutOsint(ctx, &osint.PutOsintRequest{
		ProjectId: projectID,
		Osint: &osint.OsintForUpsert{
			ProjectId:    projectID,
			ResourceType: "test",
			ResourceName: "https://example.com",
		},
	})
	handleError(err)
	osintID := putResp.Osint.OsintId
	logger.Info("Success PutOsint API.", slog.Int("osint_id", int(osintID)))

	err = client.DeleteOsint(ctx, &osint.DeleteOsintRequest{
		ProjectId: projectID,
		OsintId:   osintID,
	})
	handleError(err)
	logger.Info("Success DeleteOsint API.")
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
