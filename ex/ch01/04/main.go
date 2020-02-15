package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	appearedFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, appearedFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, appearedFiles)
			f.Close()

		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, appearedFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, appearedFiles map[string][]string) {
	input := bufio.NewScanner(f)

	line := 0
	for input.Scan() {
		line++
		counts[input.Text()]++
		appearedFiles[input.Text()] = append(
			appearedFiles[input.Text()],
			fmt.Sprintf("%s:%d", f.Name(), line),
		)
	}
}
