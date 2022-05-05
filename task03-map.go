package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0)
	for key, _ := range input {
		keys = append(keys, key)
	}

	sort.Sort(sort.IntSlice(keys))

	for _, key := range keys {
		result = append(result, input[key])
	}
	
	return
}
