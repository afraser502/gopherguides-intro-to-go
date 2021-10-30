package main

import (
	"bytes"
	"testing"
)

func Test_venue_logs(t *testing.T) {
	t.Parallel()

	bb := &bytes.Buffer{}

	v := &Venue{Log: bb}

	err := v.Entertain(1, &Band{PlayedFor: 10})

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}
