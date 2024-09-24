package exo

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}


	count := 0
	end := intervals[0][1]


	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++ 
		} else {
			end = intervals[i][1] 
		}
	}

	return count
}
