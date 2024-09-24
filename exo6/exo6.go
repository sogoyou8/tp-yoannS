package exo

func Ft_min_window(s, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	window := make(map[byte]int)
	minLen := len(s) + 1
	res := ""

	for right < len(s) {
		window[s[right]]++

		for check(window, need) && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				res = s[left : right+1]
			}
			window[s[left]]--
			left++
		}
		right++
	}

	return res
}

func check(window, need map[byte]int) bool {
	for k, v := range need {
		if window[k] < v {
			return false
		}
	}
	return true
}
