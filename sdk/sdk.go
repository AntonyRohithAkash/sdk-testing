package sdk

type SDK struct {
	Id          string
	Name        string
	Description string
}

func MyFunction(input string) string {
	return "Hello " + input

}
