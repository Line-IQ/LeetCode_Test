package funcs

func RomanToInt(s string) int {
	match := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "XL": 40, "L": 50, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if _, ok := match[s[i:i+2]]; ok {
				sum += match[s[i:i+2]]
				i++
				continue
			}
		}
		sum += match[string(s[i])]
	}
	return sum
}

func romanToInt(s string) int {
	var value = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	count := 0
	l := len(s)
	for i, _ := range s {
		if i == l-1 {
			break
		}
		if value[s[i]] < value[s[i+1]] {
			count -= value[s[i]]
		} else {
			count += value[s[i]]
		}
	}
	count += value[s[l-1]]
	return count
}
