package main

import (
	"fmt"

	"github.com/chonla/numstar"
)

func main() {
	ns := numstar.New()
	fmt.Println(ns.Big("123766"))
	fmt.Println(ns.Big("0123456789"))
}
