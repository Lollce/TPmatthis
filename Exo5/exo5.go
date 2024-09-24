package Exo

func Ft_max_substring(s string) int {
	lastIndex := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, char := range s {
		if idx, found := lastIndex[char]; found && idx >= start {
			start = idx + 1
		}
		lastIndex[char] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
