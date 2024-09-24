package exo

func Ft_max_substring(s string) int {
	seen := make(map[rune]int)
	maxLen := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if i, ok := seen[rune(s[end])]; ok && i >= start {
			start = i + 1
		}

		seen[rune(s[end])] = end

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
