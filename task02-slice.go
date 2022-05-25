package homework

func reverse(input []int64) (result []int64) {

	var reversedSlice []int64

	for _, i := range input {
		reversedSlice = append(reversedSlice, input[int64(len(input))-i])
	}

	return reversedSlice
}
