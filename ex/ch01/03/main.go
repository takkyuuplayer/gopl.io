package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func concat() {
	start := time.Now()

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())
}

func concat2() {
	start := time.Now()

	var s, sep string

	for _, val := range os.Args {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())
}

func join() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Printf("%dns elapsed\n", time.Since(start).Nanoseconds())
}

func main() {
	concat()
	concat2()
	join()
}
