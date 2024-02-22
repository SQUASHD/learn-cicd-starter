package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}{
		{"No header", nil, "", true},
		{"No Authorization header", http.Header{}, "", true},
		{"No ApiKey", http.Header{"Authorization": []string{"123456"}}, "", true},
		{"ApiKey", http.Header{"Authorization": []string{"ApiKey 123456"}}, "123456", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headerStr, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if headerStr != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", headerStr, tt.want)
			}

		})
	}

}
