package funcs

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string
	for i, n := range nums {
		j := num / n
		for k := 0; k < j; k++ {
			result += romans[i]
		}
		num = num % n
	}
	return result
}
