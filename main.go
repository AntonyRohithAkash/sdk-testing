package main

import "fmt"

type ZoneType int

const (
	Primary ZoneType = iota

	Secondary

	Alias
)

type Name struct {
	A int
}

type ErrorResponse struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorString      string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func main() {

	one := ErrorResponse{
		ErrorCode: 1,
	}
	two := ErrorResponse{
		ErrorCode: 2,
	}

	testErrors := make([]ErrorResponse, 5)
	testErrors[0] = one
	testErrors[0] = two

	//newError := newtestError.(*[]ErrorResponse)

	fmt.Printf("%v %v  ", len(testErrors), cap(testErrors))
	// commits := map[string]interface{}{
	// 	"rsc": 3711,
	// 	"r":   2138,
	// 	"gri": 1908,
	// 	"adg": 912,
	// }

	// i, ok := commits["rsc"].(string)
	// a := Name{}
	// //a.A = i.(int)
	// fmt.Println(i)
	// fmt.Println(a, ok)
}
