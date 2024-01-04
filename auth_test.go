package risken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignin(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*httptest.Server)
		wantErr bool
	}{
		{
			name: "successful signin",
			setup: func(server *httptest.Server) {
				server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.Write([]byte(`{"project_id": 1, "access_token_id": 1}`))
				})
			},
			wantErr: false,
		},
		{
			name: "failed signin",
			setup: func(server *httptest.Server) {
				server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.DefaultServeMux)
			defer server.Close()

			tt.setup(server)
			c := NewClient("test-token", WithAPIEndpoint(server.URL))
			_, err := c.Signin(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Signin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
