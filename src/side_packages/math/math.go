package math

func Average(nums []float64)float64 {
	total := float64(0)
	for _, x := range (nums) {
		total += x
	}
	if len(nums) > 1 {
		return total / float64(len(nums))

	}else{
		return total
	}
}