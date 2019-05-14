package string

// 取输入s中最长回文字符串
func longestPalindrome(s string) string {
	var ret string
	for i := 0; i < len(s); i++ {
		palindrome := ""
		for j := len(s) - 1; j >= i; j-- {
			if len(palindrome) > 0 {
				if s[j] == s[j-1] { // 重复相同的
				}
				if len(palindrome) > 1 && s[j] == palindrome[1] {
					if len(s)-j > len(palindrome) && palindrome[1:] == s[j:j+len(palindrome)-1] {
						if j+len(palindrome)-i > len(ret) {
							ret = s[i : j+len(palindrome)-1]
						}
						continue
					} else {
						palindrome = string(s[j]) + palindrome
					}

				}
			} else {
				palindrome = string(s[j]) + palindrome
				ret = palindrome
			}
		}
	}
	return ret
}

func convert(s string, numRows int) string {
	var ret string
	var a = make([][]string, 0)
	for x, v := range s {
		// a[m, n]
		if numRows == 0 {
			return ""
		}
		if numRows == 1 {
			return s
		}
		c := x / (2*numRows - 2)
		y := x % (2*numRows - 2)
		qs := c*(numRows-1) + 1

		m := qs + y/numRows*((y+1)%numRows)
		//n := y/numRows*l - (y % numRows)
		var n int
		z := x % (2*numRows - 2)
		if z+1 > numRows {
			n = 2*numRows - z - 1
		} else {
			n = z + 1
		}

		if len(a) < n {
			for i := n - len(a); i > 0; i-- {
				a = append(a, make([]string, 0))
			}
		}
		if len(a[n-1]) < m {
			for i := m - len(a[n-1]); i > 0; i-- {
				a[n-1] = append(a[n-1], "")
			}
		}
		a[n-1][m-1] = string(v)
	}
	for _, v := range a {
		for _, v1 := range v {
			ret += v1
		}
	}

	return ret
}
