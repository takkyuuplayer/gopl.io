package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)

	defer fmt.Printf("to %d\n", x)
	defer fmt.Printf("defer %d\n", x)
	defer fmt.Printf("from %d\n", x)

	f(x - 1)
}
