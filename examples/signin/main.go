package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/go-risken"
)

const (
	RISKEN_ENDPOINT = "http://localhost:8001"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("RISKEN_API_TOKEN")
	client := risken.NewClient(token, risken.WithAPIEndpoint(RISKEN_ENDPOINT))

	resp, err := client.Signin(ctx)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Success signin API.", slog.Any("resp", resp))
}
