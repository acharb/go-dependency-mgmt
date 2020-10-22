package person

import (
	"fmt"
	"test/employee"
)

type Person struct {
	Name string
}

type salary struct {
	Amount int
}

func PrintEmployee() {
	e := employee.Employee{"SDF"}
	fmt.Println(e)

	s := salary{250}
	fmt.Println(s)
}
