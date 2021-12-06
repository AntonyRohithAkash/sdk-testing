package sdk

import (
	"os"
	"testing"
)

func TestMyFunction(t *testing.T) {
	env_Variable := os.Getenv("MY_SECRET")
	expected := &SDK{ID: 1, Name: "ROHITH"}
	found := &SDK{ID: 1, Name: env_Variable}

	if env_Variable != "ROHITH" {
		t.Fatal(env_Variable)
	}

	if expected.String() != found.String() {
		t.Fatalf("Expected:%s\nFound:%s", expected, found)
	}
}
