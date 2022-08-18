package methodexample

import "fmt"

type Person struct {
	Fullname string
	Age      int
}

func (p Person) PersonLogger() string {
	return fmt.Sprintf("%s age %d", p.Fullname, p.Age)
}
