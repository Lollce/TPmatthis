package Exo

func Ft_missing(nums []int) int {
	max := 0
	min := 0

	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if min >= nums[i] {
			min = nums[i] + 1
		}
	}

	if min < 0 {
		return 0
	} else {
		for cpt := min; cpt <= max; cpt++ {

			found := false

			for a := 0; a < len(nums); a++ {
				if cpt == nums[a] {
					found = true
					break
				}
			}

			if found != true {
				return cpt
			}
		}
		return max + 1
	}
}
