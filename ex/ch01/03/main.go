package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	from := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(time.Since(from).Nanoseconds(), "ns elapsed.")

	from = time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(from).Nanoseconds(), "ns elapsed.")
}
