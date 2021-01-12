package comma

import (
	"fmt"
	"math/rand"
	"testing"
)

var randoNum = rand.Intn(9999999)
var randoString = fmt.Sprint(randoNum)

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma(randoString)
	}
}

func TestComma(t *testing.T) {
	if comma("123") != "123" {
		t.Error("123 != 123")
	}
	if comma("1234") != "1,234" {
		t.Error("1234 != 1,234")
	}
}
