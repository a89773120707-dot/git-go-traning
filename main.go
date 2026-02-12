package main

import (
	"fmt"
	"strings"
	"sync"
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

func double(n int) int {
	return n * 2
}

func UpperCheckAsci(name string) bool {
	if name == "" {
		return false
	}

	first := name[0]

	return first >= 'A' && first <= 'Z'

}

// fan in
func gen1(ch chan<- int) {
	for i := 1; i <= 4; i++ {
		ch <- i*10 + 1
	}
}

func gen2(ch chan<- int) {
	for i := 1; i <= 4; i++ {
		ch <- i*100 + 1
	}
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	fmt.Println(Greet(""))
	fmt.Println(Greet("anar"))
	fmt.Println(Greet("Anar"))

	fmt.Println(Square(2))
	fmt.Println(Square(4))

	fmt.Println("Isvalidname('Alex'):", IsValidName("Alex"))

	fmt.Println(double(7))

	fmt.Println("UpperCheckAsci name anar - ", UpperCheckAsci("anar"))
	fmt.Println("UpperCheckAsci name Anar - ", UpperCheckAsci("Anar"))

	//fan in
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		gen1(ch1)
		close(ch1)
	}()

	go func() {
		gen2(ch2)
		close(ch2)
	}()

	merged := merge(ch1, ch2)

	for v := range merged {
		fmt.Println("Получили из merged :", v)
	}

	fmt.Println("Все ...")

}
