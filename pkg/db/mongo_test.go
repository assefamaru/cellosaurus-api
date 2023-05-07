package db

import (
	"errors"
	"testing"
)

func TestNewMongoConfigFromEnv(t *testing.T) {
	testCases := []struct {
		name    string
		host    string
		port    string
		wantErr error
	}{
		{
			"Valid input",
			"gotham",
			"comic",
			nil,
		},
		{
			"Missing host",
			"",
			"comic",
			errMissingEnv,
		},
		{
			"Missing port",
			"gotham",
			"",
			errMissingEnv,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(mongoServiceHostEnv, tc.host)
			t.Setenv(mongoServicePortEnv, tc.port)

			_, err := newMongoConfigFromEnv()

			if got, want := err, tc.wantErr; !errors.Is(got, want) {
				t.Errorf("newMongoConfigFromEnv(%s): got = '%v', want = '%v'", tc.name, got, want)
			}
		})
	}
}
