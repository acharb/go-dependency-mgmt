package main

import (
	"fmt"

	"github.com/stellar/go/support/log"
)

func main() {

	b := Boss{Position: "manager"}
	fmt.Println(b)

	log.Info("hi")

}

type Boss struct {
	Position   string
	Department string
}
