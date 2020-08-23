package math

func power(base int, exponent int) int {
	res := 1
	for x := 0; x < exponent; x++ {
		res *= base
	}
	return res
}
