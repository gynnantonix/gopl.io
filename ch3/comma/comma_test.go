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
	if ret := comma("123"); ret != "123" {
		t.Error("comma(123): " + ret)
	}
	if ret := comma("1234"); ret != "1,234" {
		t.Error("comma(1234): " + ret)
	}
	if ret := comma("-1234"); ret != "-1,234" {
		t.Error("comma(-1234): " + ret)
	}
	if ret := comma("-1234.5678"); ret != "-1,234.5678" {
		t.Error("comma(-1234.5678): " + ret)
	}
}
