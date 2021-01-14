package main

func main() {

}

// reverses the order of x in-place
func reverse(x []int) {
	for i, j := 0, len(x); i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}

// rotates x left by n elements
func lRotN(x []int, n int) {
	reverse(x[n:])
	reverse(x[:n])
	reverse(x)
}

// rotates x right by n elements
func rRotN(x []int, n int) {
	reverse(x)
	reverse(x[n:])
	reverse(x[:n])
}
