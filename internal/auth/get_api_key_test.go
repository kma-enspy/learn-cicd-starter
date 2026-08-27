package auth

import (
    "reflect"
    "testing"
	"error"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string][]string
		want    string
		wantErr error
	}{
		{
			name: "valid api key",
			headers: map[string][]string{
				"Authorization": {"ApiKey my-api-key"},
			},
			want:    "my-api-key",
			wantErr: nil,
		},
		{
			name: "missing authorization header",
			headers: map[string][]string{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			headers: map[string][]string{
				"Authorization": {"Bearer my-api-key"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
