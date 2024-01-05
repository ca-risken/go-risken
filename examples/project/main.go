package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/core/proto/project"
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
	logger.Info("Success Signin.", slog.Int("project_id", int(projectID)))

	listResp, err := client.ListProject(ctx, &project.ListProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success ListProject API.", slog.Int("len", len(listResp.Project)))

	updateResp, err := client.UpdateProject(ctx, &project.UpdateProjectRequest{
		ProjectId: projectID,
		Name:      listResp.Project[0].Name,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success UpdateProject API.", slog.String("name", updateResp.Project.Name))

	tagResp, err := client.TagProject(ctx, &project.TagProjectRequest{
		ProjectId: projectID,
		Tag:       "test",
		Color:     "red",
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success ListTag API.", slog.Any("len", tagResp.ProjectTag.Tag))

	if err := client.UntagProject(ctx, &project.UntagProjectRequest{
		ProjectId: projectID,
		Tag:       "test",
	}); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success UntagProject API.")
}
