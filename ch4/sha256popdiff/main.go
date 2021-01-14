package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	testSha := sha256.Sum256([]byte(" "))
	testSha2 := sha256.Sum256([]byte("0"))
	fmt.Printf("[%s]: %x, %d\n", []byte(" "), testSha, popCount(testSha))
	fmt.Printf("[%s]: %x, %d\n", []byte(" "), testSha2, popCount(testSha2))
	fmt.Printf("%d\n", int(math.Abs(float64(popCount(testSha)-popCount(testSha2)))))
}

// PopCount returns the population count (number of set bits) of a byte array
func popCount(x [32]byte) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += int(pc[x[i]])
	}
	return sum
}
