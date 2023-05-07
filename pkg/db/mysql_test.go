package db

import (
	"errors"
	"testing"
)

func TestNewMySQLConfigFromEnv(t *testing.T) {
	testCases := []struct {
		name    string
		user    string
		pass    string
		db      string
		host    string
		port    string
		wantErr error
	}{
		{
			"Valid input",
			"brucewayne",
			"batman",
			"justiceleague",
			"gotham",
			"comic",
			nil,
		},
		{
			"Missing user",
			"",
			"batman",
			"justiceleague",
			"gotham",
			"comic",
			errMissingEnv,
		},
		{
			"Missing pass",
			"brucewayne",
			"",
			"justiceleague",
			"gotham",
			"comic",
			errMissingEnv,
		},
		{
			"Missing db",
			"brucewayne",
			"batman",
			"",
			"gotham",
			"comic",
			errMissingEnv,
		},
		{
			"Missing host",
			"brucewayne",
			"batman",
			"justiceleague",
			"",
			"comic",
			errMissingEnv,
		},
		{
			"Missing port",
			"brucewayne",
			"batman",
			"justiceleague",
			"gotham",
			"",
			errMissingEnv,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(mysqlServiceUserEnv, tc.user)
			t.Setenv(mysqlServicePassEnv, tc.pass)
			t.Setenv(mysqlServiceDBEnv, tc.db)
			t.Setenv(mysqlServiceHostEnv, tc.host)
			t.Setenv(mysqlServicePortEnv, tc.port)

			_, err := newMySQLConfigFromEnv()

			if got, want := err, tc.wantErr; !errors.Is(got, want) {
				t.Errorf("newMySQLConfigFromEnv(%s): got = '%v', want = '%v'", tc.name, got, want)
			}
		})
	}
}
