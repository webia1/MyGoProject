package main

import (
	"errors"
	"fmt"
)

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf(" resource %s: code %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {

	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}

	return false
}

func main() {

	err1 := ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     456,
	}

	if errors.Is(err1, ResourceErr{Resource: "Database"}) {
		fmt.Println("err1:", err1)
		// err1:  resource Database: code 123
	}
	if errors.Is(err1, ResourceErr{Code: 123}) {
		fmt.Println("err1:", err1)
		// err1:  resource Database: code 123
	}

	if errors.Is(err2, ResourceErr{Resource: "Network"}) {
		fmt.Println("err2:", err2)
		// err2:  resource Network: code 456
	}

	if errors.Is(err2, ResourceErr{Code: 456}) {
		fmt.Println("err2:", err2)
		// err2:  resource Network: code 456
	}

	if errors.Is(err1, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("err1:", err1)
		// err1:  resource Database: code 123
	}

	if errors.Is(err2, ResourceErr{Resource: "Network", Code: 456}) {
		fmt.Println("err2:", err2)
		// err2:  resource Network: code 456
	}

	if errors.Is(err1, ResourceErr{Resource: "Hohoho"}) {
		fmt.Println("err1:", err1)
		// No output
	}

	if errors.Is(err1, ResourceErr{Resource: "Hohoho", Code: 123}) {
		fmt.Println("err1:", err1)
		// No output
	}

	if errors.Is(err2, ResourceErr{Resource: "Hohoho"}) {
		fmt.Println("err2:", err2)
		// No output
	}

	if errors.Is(err2, ResourceErr{Resource: "Hohoho", Code: 456}) {
		fmt.Println("err2:", err2)
		// No output
	}

	fmt.Println("Debugger")
}

/*

 */
