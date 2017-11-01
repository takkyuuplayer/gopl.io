package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func wordfreq(in io.Reader) {
	counts := make(map[string]int)

	input := bufio.NewScanner(in)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	for word, n := range counts {
		fmt.Printf("%q\t%d\n", word, n)
	}
}

func main() {
	wordfreq(strings.NewReader("This is a test. That is also a test"))

	wordfreq(os.Stdin)
}
