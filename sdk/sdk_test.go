package sdk

import "testing"

func TestMyFunction(t *testing.T) {
	input := "Rohith"
	expected := "Hello !Rohith"
	actual := MyFunction(input)

	if expected != actual {
		t.Fatal("Data Mismatched")
	}
}
