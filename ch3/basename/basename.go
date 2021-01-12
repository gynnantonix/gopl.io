// Package basename does what the program of the same name in UNIX does
package basename

// Basename1 does it without using any libraries
func Basename1(s string) string {
	// Discard last '/' and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// preserve everything before the last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
