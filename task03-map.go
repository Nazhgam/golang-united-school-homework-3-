package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	intMap := []int{}
	for k, _ := range input {
		intMap = append(intMap, k)
	}
	sort.Ints(intMap)
	for i := 0; i < len(intMap); i++ {
		result = append(result, input[intMap[i]])
	}
	return
}
