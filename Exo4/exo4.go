package Exo

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxbenef := 0
	minPrix := prices[0]

	for _, price := range prices[1:] {
		if price < minPrix {
			minPrix = price
		} else {
			profit := price - minPrix
			if profit > maxbenef {
				maxbenef = profit
			}
		}
	}

	return maxbenef
}
