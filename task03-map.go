package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var arrayKeys = make([]int, 0)
	for key, _ := range input {
		arrayKeys = append(arrayKeys, key)
	}
	sort.Ints(arrayKeys)
	for _, key := range arrayKeys {
		for keyInput, valueInput := range input {
			if key == keyInput {
				result = append(result, valueInput)
				break
			}
		}
	}
	return result
}
