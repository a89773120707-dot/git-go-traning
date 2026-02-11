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
func Square(n int) int {
	return n * n
}
func main() {
	fmt.Println(Greet(""))
	fmt.Println(Greet("anar"))
	fmt.Println(Greet("Anar"))

	fmt.Println(Square(2))
	fmt.Println(Square(4))

}
