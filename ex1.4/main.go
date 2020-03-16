package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printData(name string, counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %s\t%d\n", name, line, count)
		}
	}
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		printData("os.Stdin", counts)
	} else {
		counts := make(map[string]int)
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			printData(arg, counts)
		}
	}
}
