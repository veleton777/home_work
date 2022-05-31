package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	baseText := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(baseText))
}
