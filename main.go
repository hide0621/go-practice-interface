package main

import (
	"fmt"
	"go-practice-interface/counterutil" //モジュール名/パッケージ名でインポート
)

func main() {
	var b counterutil.ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
