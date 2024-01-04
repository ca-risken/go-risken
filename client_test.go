package risken

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testClient = NewClient("token", WithAPIEndpoint("http://localhost"))

func TestCheckResponse(t *testing.T) {
	type args struct {
		resp *httptest.ResponseRecorder
		err  error
	}
	tests := []struct {
		name    string
		args    args
		want    *httptest.ResponseRecorder
		wantErr error
	}{
		{
			name: "200 OK",
			args: args{
				resp: &httptest.ResponseRecorder{Code: 200},
				err:  nil,
			},
			want:    &httptest.ResponseRecorder{Code: 200},
			wantErr: nil,
		},
		{
			name: "500 Error",
			args: args{
				resp: &httptest.ResponseRecorder{Code: 500},
				err:  errors.New("500 Error"),
			},
			want:    nil,
			wantErr: errors.New("error calling the API endpoint: 500 Error"),
		},
		{
			name: "404 Error(Unexpected Content-Type)",
			args: args{
				resp: &httptest.ResponseRecorder{
					Code: 404,
					Body: bytes.NewBuffer([]byte(`{"status":404,"error":"404 Not Found"}`)),
				},
				err: nil,
			},
			want:    nil,
			wantErr: errors.New("HTTP response with status code 404 does not contain Content-Type: application/json"),
		},
		{
			name: "404 Error(API Error)",
			args: args{
				resp: &httptest.ResponseRecorder{
					Code: 404,
					HeaderMap: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: bytes.NewBuffer([]byte(`{"error":"404 Not Found"}`)),
				},
				err: nil,
			},
			want:    nil,
			wantErr: APIError{Status: 404, Message: "404 Not Found"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpResponse := tt.args.resp.Result()

			resp, err := testClient.checkResponse(httpResponse, tt.args.err)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}

			var wantReponse *http.Response
			if tt.want != nil {
				wantReponse = tt.want.Result()
			}
			if diff := cmp.Diff(resp, wantReponse); diff != "" {
				t.Errorf("checkResponse() error, diff=%s", diff)
			}
		})
	}
}
