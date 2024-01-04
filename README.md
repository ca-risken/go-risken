# go-risken

go client library for RISKEN API

## Client Library

`RISKEN_API_TOKEN` is required before using the library. 
Please create a token on RISKEN beforehand. 
See [document](https://docs.security-hub.jp/risken/access_token/) for details.

### Example

```go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ca-risken/go-risken"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	ctx := context.Background()
	token := os.Getenv("RISKEN_API_TOKEN") // Retrieve api token from environment value.
	client := risken.NewClient(token, risken.WithAPIEndpoint("http://your-risken-api-endpoint"))
	resp, err := client.Signin(ctx)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Success signin API.", slog.Any("resp", resp))
}
```
