package homework

func average(input [15]float32) (result float32) {
	elements := float32(len(input))
	result = 0
	for _, item := range input {
		result += item
	}
	return result / elements
}
