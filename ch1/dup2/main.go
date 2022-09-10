package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filename := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "", filename)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, filename)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filename[line], " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, arg string, filename map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if isHere(filename[input.Text()], arg) {
			filename[input.Text()] = append(filename[input.Text()], arg)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func isHere(files []string, file string) bool {
	for _, f := range files {
		if f == file {
			return false
		}
	}
	return true
}
