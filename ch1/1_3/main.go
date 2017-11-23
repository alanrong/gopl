package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
	end := time.Now().UnixNano()
	diff := end - start
	fmt.Println("inefficient timecost:%d", diff)

	start = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		_ = strings.Join(os.Args[1:], " ")
	}
	end = time.Now().UnixNano()
	diff = end - start
	fmt.Println("strings join timecost:%d", diff)

}
