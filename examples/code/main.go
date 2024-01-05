package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/datasource-api/proto/code"
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

	listResp, err := client.ListGitHubSetting(ctx, &code.ListGitHubSettingRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListGitHubSetting API.", slog.Int("len", len(listResp.GithubSetting)))

	putResp, err := client.PutGitHubSetting(ctx, &code.PutGitHubSettingRequest{
		ProjectId: projectID,
		GithubSetting: &code.GitHubSettingForUpsert{
			ProjectId:           projectID,
			Name:                "test",
			GithubUser:          "hoge",
			PersonalAccessToken: "dummy",
			Type:                code.Type_ORGANIZATION,
			TargetResource:      "ca-risken/go-risken",
		},
	})
	handleError(err)
	githubID := putResp.GithubSetting.GithubSettingId
	logger.Info("Success PutGitHubSetting API.", slog.Int("github_setting_id", int(githubID)))

	putCodeScanResp, err := client.PutCodeScanSetting(ctx, &code.PutCodeScanSettingRequest{
		ProjectId: projectID,
		CodeScanSetting: &code.CodeScanSettingForUpsert{
			ProjectId:        projectID,
			GithubSettingId:  githubID,
			CodeDataSourceId: 1,
			Status:           code.Status_CONFIGURED,
		},
	})
	handleError(err)
	logger.Info("Success PutCodeScanSetting API.", slog.Any("resp", putCodeScanResp))

	err = client.DeleteGitHubSetting(ctx, &code.DeleteGitHubSettingRequest{
		ProjectId:       projectID,
		GithubSettingId: githubID,
	})
	handleError(err)
	logger.Info("Success DeleteGitHubSetting API.")
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
