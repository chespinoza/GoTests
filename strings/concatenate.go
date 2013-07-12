package strings

import "bytes"

// concatenate arrays
func stringConcatenate(x, y []string) []string {
	result := make([]string, len(x)+len(y))
	i := 0
	for ; i < len(x); i++ {
		result[i] = x[i]
	}
	for j := 0; i < len(x)+len(y); i++ {
		result[i] = y[j]
		j++
	}
	return result
}

// more idiomatic concatenate slices
func concat(a, b []string) []string {
	c := make([]string, len(a)+len(b))
	copy(c[:len(a)], a)
	copy(c[len(a):], b)
	return c
}

// adapted to a variadic way from http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func conc(str ...string) string {
	var buffer bytes.Buffer
	for _, val := range str {
		buffer.WriteString(val)
	}
	return buffer.String()
}
