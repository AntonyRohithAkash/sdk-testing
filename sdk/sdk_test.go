package sdk

import (
	"os"
	"testing"
)

func TestMyFunction(t *testing.T) {
	envVariable := os.Getenv("MY_SECRET")
	expected := &SDK{Id: 1, Name: "ROHITH"}
	found := &SDK{Id: 1, Name: envVariable}

	if envVariable != "ROHITH" {
		t.Fatal(envVariable)
	}

	if expected.String() != found.String() {
		t.Fatalf("Expected:%s\nFound:%s", expected, found)
	}
}
