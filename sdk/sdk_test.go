package sdk

import (
	"testing"
)

func TestMyFunction(t *testing.T) {
	expected := "ID:1 - Name:Rohith"
	sdk := SDK{ID: 1, Name: "Rohith"}
	found := sdk.String()
	t.Fatalf("mismatch")
	if found != expected {
		t.Fatalf("mismatch")
	}
}
