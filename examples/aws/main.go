package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/datasource-api/proto/aws"
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

	listResp, err := client.ListAWS(ctx, &aws.ListAWSRequest{
		ProjectId: projectID,
	})
	handleError(err)
	logger.Info("Success ListAWS API.", slog.Int("len", len(listResp.Aws)))

	putResp, err := client.PutAWS(ctx, &aws.PutAWSRequest{
		ProjectId: projectID,
		Aws: &aws.AWSForUpsert{
			ProjectId:    projectID,
			Name:         "test",
			AwsAccountId: "123456789012", // dummy account id
		},
	})
	handleError(err)
	awsID := putResp.Aws.AwsId
	testAwsDataSourceID := uint32(1)
	logger.Info("Success PutAWS API.", slog.Int("aws_id", int(awsID)))

	attachResp, err := client.AttachDataSource(ctx, &aws.AttachDataSourceRequest{
		ProjectId: projectID,
		AttachDataSource: &aws.DataSourceForAttach{
			AwsId:           awsID,
			AwsDataSourceId: testAwsDataSourceID,
			ProjectId:       projectID,
			AssumeRoleArn:   "arn:aws:iam::123456789012:role/risken", // dummy role arn
			ExternalId:      "dummydummydummydummydummy",
			Status:          aws.Status_CONFIGURED,
		},
	})
	handleError(err)
	logger.Info("Success AttachDataSource API.", slog.Any("aws_data_source", attachResp))

	err = client.DetachDataSource(ctx, &aws.DetachDataSourceRequest{
		ProjectId:       projectID,
		AwsId:           awsID,
		AwsDataSourceId: testAwsDataSourceID,
	})
	handleError(err)
	logger.Info("Success DetachDataSource API.")

	err = client.DeleteAWS(ctx, &aws.DeleteAWSRequest{
		ProjectId: projectID,
		AwsId:     awsID,
	})
	handleError(err)
	logger.Info("Success DeleteAWS API.")
}

func handleError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
