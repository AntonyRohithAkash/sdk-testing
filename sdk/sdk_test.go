package sdk

import (
	"os"
	"testing"
)

func TestMyFunction(t *testing.T) {
	input := "Rohith"
	expected := "Hello Rohith"
	actual := MyFunction(input)
	
	env_varaiable  := os.Getenv("MY_SECRET")
	if env_varaiable !="ROHIT"{
		t.Fatal(env_varaiable)
	}

	if expected != actual {
		t.Fatal("Data Mismatched")
	}
}
