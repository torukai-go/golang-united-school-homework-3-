package homework

func reverse(input []int64) (result []int64) {

	var reversedSlice []int64

	for i := len(input); i > 0; i-- {
		reversedSlice = append(reversedSlice, input[i])
	}

	return reversedSlice
}
