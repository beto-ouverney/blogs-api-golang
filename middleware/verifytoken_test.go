package middleware

import (
	"encoding/json"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	type args struct {
		header map[string][]string
		body   *json.Decoder
	}
	tests := []struct {
		name        string
		args        args
		wantOk      bool
		wantStatus  int
		wantMessage string
	}{
		{
			name: "Test if a valid token is valid",
			args: args{
				header: map[string][]string{
					"Authorization": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaXNwbGF5TmFtZSI6IkJyZXR0IFdpbHRzaGlyZSIsImVtYWlsIjoiYnJldHRAZW1haWwuY29tIiwiZXhwIjoxNjU4NjM3NDU0LCJpbWFnZSI6Imh0dHA6Ly80LmJwLmJsb2dzcG90LmNvbS9fWUE1MGFkUS03dlEvUzFnZlJfNnVmcEkvQUFBQUFBQUFBQWsvMUVySkdnUldaRGcvUzQ1L2JyZXR0LnBuZyJ9.hdhlI8CuG9xOioNeaN4fh4cy2SH_5S_zi4cwRVMyL88"},
				},
				body: &json.Decoder{},
			},
			wantOk:      true,
			wantStatus:  200,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotStatus, gotMessage := VerifyToken(tt.args.header, tt.args.body)
			if gotOk != tt.wantOk {
				t.Errorf("VerifyToken() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("VerifyToken() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("VerifyToken() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
