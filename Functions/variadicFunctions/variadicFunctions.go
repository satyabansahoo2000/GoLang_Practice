package variadicFunctions

func AddNums(nums ...int) (total int) {
	total = 0
	for _, n := range nums {
		total += n
	}
	return 
}