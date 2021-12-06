package sdk

import "fmt"

type SDK struct {
	ID   int
	Name string
}

func (sdk *SDK) String() string {
	return fmt.Sprintf("ID:%v - Name:%v", sdk.ID, sdk.Name)
}
