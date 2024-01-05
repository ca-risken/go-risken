package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/datasource-api/proto/google"
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

	listResp, err := client.ListGCP(ctx, &google.ListGCPRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListGCP API.", slog.Int("len", len(listResp.Gcp)))

	putResp, err := client.PutGCP(ctx, &google.PutGCPRequest{
		ProjectId: projectID,
		Gcp: &google.GCPForUpsert{
			ProjectId:        projectID,
			Name:             "test",
			GcpProjectId:     "test",
			VerificationCode: "dummydummydummy",
		},
	})
	handleError(err)
	gcpID := putResp.Gcp.GcpId
	logger.Info("Success PutGCP API.", slog.Int("gcp_id", int(gcpID)))

	err = client.DeleteGCP(ctx, &google.DeleteGCPRequest{
		ProjectId: projectID,
		GcpId:     gcpID,
	})
	handleError(err)
	logger.Info("Success DeleteGCP API.")
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
