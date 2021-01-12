// Package comma puts commas in numbers
package comma

import "strings"

// comma trims whitespace, sign, and floating-point component, and
// calls commaInner
func comma(s string) string {
	sign := ""
	fp := ""

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return s
	}

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}
	if s[0] == '-' || s[0] == '+' {
		sign = string(s[0])
		s = s[1:]
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		fp = s[dot:]
		s = s[:dot]
	}
	return sign + commaInner(s) + fp
}

// commaInner inserts commas in a non-negative decimal integer string
func commaInner(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaInner(s[:n-3]) + "," + s[n-3:]
}
