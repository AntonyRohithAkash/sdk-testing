package sdk

import "testing"

func TestMyFunction(t *testing.T) {
	input := "Rohith"
	expected := "Hello Rohith"
	actual := MyFunction(input)
	
	env_varaiable  = os.Getenv("MY_SECRET")
	if env_varaiable !="NEW"{
		t.Fatal(env_varaiable)
	}

	if expected != actual {
		t.Fatal("Data Mismatched")
	}
}
