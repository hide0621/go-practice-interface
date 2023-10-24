package main

import (
	"fmt"
	"go-practice-interface/countertil"
)

func main() {
	var b countertil.ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
