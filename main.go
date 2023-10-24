package main

import "fmt"

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc = ByteCounter(len(p))
	return len(p), err
}

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
