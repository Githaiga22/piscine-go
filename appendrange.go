package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var array []int
	for i := min; i < max; i++ {
		array = append(array, i)
	}
	return array
}
