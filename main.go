package main

import (
	"fmt"
	"strings"
	"unicode"
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

func IsValidName(name string) bool {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return false
	}
	if len(trimmed) < 2 || len(trimmed) > 50 {
		return false
	}

	for _, r := range trimmed {
		if !unicode.IsLetter(r) && r != ' ' && r != '-' && r != '\'' {
			return false
		}

	}
	return true
}

func main() {
	fmt.Println(Greet(""))
	fmt.Println(Greet("anar"))
	fmt.Println(Greet("Anar"))

	fmt.Println(Square(2))
	fmt.Println(Square(4))

	fmt.Println("Isvalidname('Alex'):", IsValidName("Alex"))

}
