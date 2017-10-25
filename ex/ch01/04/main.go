package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	appears := make(map[string]map[string]int)

	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		input := bufio.NewScanner(f)

		for input.Scan() {
			counts[input.Text()]++
			if appears[input.Text()] == nil {
				appears[input.Text()] = make(map[string]int)
			}
			appears[input.Text()][filename]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)

			for filename, c := range appears[line] {
				fmt.Printf("\t%d\t%s\n", c, filename)
			}
		}
	}
}
