package funcs

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]string, len(s))
	}
	mode := 0
	j := 0
	index := 0
	for i := 0; i < numRows; index++ {
		if i == 0 {
			mode = 0
		}
		if i == numRows-1 {
			mode = 1
		}

		res[i][j] = string(s[index])

		if index == len(s)-1 {
			break
		}

		switch mode {
		case 0:
			i++
		case 1:
			i--
			j++
		}
	}
	result := consist(res)
	return result
}

func consist(s [][]string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] != "" {
				res += string(s[i][j])
			}
		}
	}
	return res
}

func BestConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n, l := len(s), 2*(numRows-1)
	chs, p := make([]byte, n), 0
	for i := 0; i < n; i, p = i+l, p+1 {
		chs[p] = s[i]
	}
	for i := 1; i < numRows-1; i++ {
		for j := i; p < n && j < n; j += l {
			chs[p] = s[j]
			p++
			mid := j + l - 2*i
			if mid < n {
				chs[p] = s[mid]
				p++
			}
		}
	}
	for i := numRows - 1; i < n; i, p = i+2*(numRows-1), p+1 {
		chs[p] = s[i]
	}
	return string(chs)
}
