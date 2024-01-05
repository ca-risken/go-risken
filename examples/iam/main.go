package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/core/proto/iam"
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

	listResp, err := client.ListUser(ctx, &iam.ListUserRequest{
		Activated: true,
		ProjectId: projectID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success ListUser API.", slog.Int("len", len(listResp.UserId)))

	if len(listResp.UserId) < 1 {
		logger.Info("No user found.")
		os.Exit(0)
	}
	userID := listResp.UserId[0]
	getResp, err := client.GetUser(ctx, &iam.GetUserRequest{
		UserId: userID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success GetUser API.", slog.String("name", getResp.User.Name))

	listRoleResp, err := client.ListRole(ctx, &iam.ListRoleRequest{
		ProjectId: projectID,
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success ListRole API.", slog.Int("len", len(listRoleResp.RoleId)))
	roleID := listRoleResp.RoleId[0]

	putUserReservedResp, err := client.PutUserReserved(ctx, &iam.PutUserReservedRequest{
		ProjectId: projectID,
		UserReserved: &iam.UserReservedForUpsert{
			UserIdpKey: "alice@example.com",
			RoleId:     roleID,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success PutUserReserved API.", slog.Any("resp", putUserReservedResp))

	if err = client.DeleteUserReserved(ctx, &iam.DeleteUserReservedRequest{
		ProjectId:  projectID,
		ReservedId: putUserReservedResp.UserReserved.ReservedId,
	}); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success DeleteUserReserved API.", slog.Any("ReservedId", putUserReservedResp.UserReserved.ReservedId))
}
