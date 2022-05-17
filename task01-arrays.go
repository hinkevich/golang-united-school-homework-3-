package homework

func average(input []float32) (result float32) {
	for _, v := range input {
		result += v
	}
	return result / float32(len(input))
}
