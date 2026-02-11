package main

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	if name == "" {
		return "Hello no name"
	}
	name = strings.Title(strings.ToLower(name))

	return "Hello " + name + "!"
}
func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(Greet(""))
	fmt.Println(Greet("anar"))
	fmt.Println(Greet("Anar"))

	fmt.Println(IsEven(2))
	fmt.Println(IsEven(3))

}
