package sdk

import "fmt"

type SDK struct {
	Id   int
	Name string
}

func (sdk *SDK) String() string {
	return fmt.Sprintf("ID:%v - Name:%v", sdk.Id, sdk.Name)
}
