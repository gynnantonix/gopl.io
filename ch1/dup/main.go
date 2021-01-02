package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string]string)
	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if !strings.Contains(files[line], arg) {
				if len(files[line]) != 0 {
					files[line] += ","
				}
				files[line] += arg
			}
		}
	}

	for line, n := range counts {
		if n == 0 {
			continue
		}
		fmt.Printf("%d\t[%-40s]\t%s\n", n, files[line], line)
	}
}
