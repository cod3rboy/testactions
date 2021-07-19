package main

import (
	"fmt"
	"stack/stack"
)


func main() {
	s := stack.New()
	s.Push("World!")
	s.Push("Hello,")
	var element1, _ = s.Pop()
	var element2, _ = s.Pop()
	fmt.Println(element1, element2)
}