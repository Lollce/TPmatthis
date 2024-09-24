package Exo

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	required := make(map[rune]int)
	for _, char := range t {
		required[char]++
	}

	left, right, formed := 0, 0, 0
	windowCounts := make(map[rune]int)
	minLength, minLeft := len(s)+1, 0
	requiredLength := len(required)

	for right < len(s) {
		c := rune(s[right])
		windowCounts[c]++

		if requiredCount, found := required[c]; found && windowCounts[c] == requiredCount {
			formed++
		}

		for left <= right && formed == requiredLength {
			c = rune(s[left])

			if right-left+1 < minLength {
				minLength = right - left + 1
				minLeft = left
			}

			windowCounts[c]--
			if requiredCount, found := required[c]; found && windowCounts[c] < requiredCount {
				formed--
			}
			left++
		}
		right++
	}

	if minLength == len(s)+1 {
		return ""
	}

	return s[minLeft : minLeft+minLength]
}
