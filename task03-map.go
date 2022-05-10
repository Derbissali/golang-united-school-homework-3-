package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	SortInt := make([]int, 0, len(input))
	for k := range input {
		SortInt = append(SortInt, k)
	}
	sort.Ints(SortInt)
	for _, i := range SortInt {
		result = append(result, input[i])

	}
	return
}
