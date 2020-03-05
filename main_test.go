package main

import (
	"errors"
	"testing"
)

// We need to be able to connect to a DB
// using a connection string and get
// a testquery result
//
func TestConnect(t *testing.T) {
	cases := []struct {
		caseName           string
		errored            bool
		error              error
		inConnectionString string
		wantResultCount    int
	}{
		{
			"connection-failure-case",
			true,
			errors.New("dial tcp 127.0.0.1:5432: connect: connection refused"),
			"user=pqgotest dbname=pqgotest sslmode=verify-full",
			0},
	}
	for _, c := range cases {
		resultCount, err := connect(c.inConnectionString)
		if err != nil {
			if c.errored != true {
				t.Errorf("case %s unexpected error (%q) ", c.caseName, err)
			} else {
				if c.error.Error() != err.Error() {
					t.Errorf("case %s expected error (%q) found (%q) ", c.caseName, c.error, err)
				}
			}
		}
		if resultCount != c.wantResultCount {
			t.Errorf("connect (%q) == %q, want %q", c.inConnectionString, resultCount, c.wantResultCount)
		}
	}
}
