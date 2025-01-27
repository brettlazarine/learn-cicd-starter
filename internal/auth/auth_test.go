package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	validHeader := http.Header{
		"Authorization": []string{"ApiKey validApiKey"},
	}
	invalidHeader := http.Header{
		"Authorization": []string{""},
	}
	malformedHeader := http.Header{
		"Authorization": []string{"not so goo"},
	}

	tests := []struct {
		name       string
		header     http.Header
		wantHeader string
		wantErr    bool
	}{
		{
			name:       "Valid header",
			header:     validHeader,
			wantHeader: "validApiKey",
			wantErr:    false,
		},
		{
			name:       "Invalid header",
			header:     invalidHeader,
			wantHeader: "",
			wantErr:    true,
		},
		{
			name:       "Malformed header",
			header:     malformedHeader,
			wantHeader: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApiKey, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApiKey != tt.wantHeader {
				t.Errorf("GetAPIKey() = %v, want %v", gotApiKey, tt.wantHeader)
			}
		})
	}
}
