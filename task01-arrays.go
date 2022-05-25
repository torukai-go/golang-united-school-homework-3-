package main

func average(input [15]float32) (result float32) {
	var a float32

	for _, s := range input {
		a = a + s
	}

	return a / float32(len(input))
}
