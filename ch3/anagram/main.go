// tests whether two strings are anagrams of each other
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing string arguments")
	}

	str1 := []byte(os.Args[1])
	str2 := []byte(os.Args[2])

	if len(str1) != len(str2) {
		fmt.Print("These are not anagrams\n")
		return
	}

	counts := make(map[byte]int)

	for idx := 0; idx < len(str1); idx++ {
		counts[str1[idx]]++
		counts[str2[idx]]--
	}
	var accum int
	for _, sum := range counts {
		accum += sum
	}

	if accum != 0 {
		fmt.Print("These are not anangrams\n")
		return
	}
	fmt.Print("These are anangrams\n")
}
