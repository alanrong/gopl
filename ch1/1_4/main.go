package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filenames := make(map[string][]string)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, arg, counts, filenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			v := filenames[line]
			for i := 0; i < len(v); i++ {
				fmt.Printf("%s\t", v[i])
			}
			fmt.Println("")
		}
	}
}

func countLines(f *os.File, arg string, counts map[string]int, files map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if v, ok := files[line]; ok {
			var flag bool
			for i := 0; i < len(v); i++ {
				if v[i] == arg {
					flag = true
					break
				}
			}
			if !flag {
				files[line] = append(v, arg)
			}
		} else {
			files[line] = []string{arg}
		}
	}
}
