package Exo

func Ft_coin(coins []int, amount int) int {

	cpt := 0

	for i := 0; i < len(coins)-1; i++ {
		for j := i + 1; j < len(coins); j++ {
			if coins[i] < coins[j] {
				coins[i], coins[j] = coins[j], coins[i]
			}
		}
	}

	for i := 0; i < len(coins); i++ {
		count := amount / coins[i]
		cpt += count
		amount -= count * coins[i]
	}

	if amount == 0 {
		return cpt
	} else {
		return -1
	}
}