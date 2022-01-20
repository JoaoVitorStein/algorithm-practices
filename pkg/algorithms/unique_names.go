package algorithms

func UniqueNames(a, b []string) []string {
	result := make([]string, 0)
	added := make(map[string]bool, 0)
	for i := 0; i < len(b); i++ {
		aVal := a[i]
		bVal := b[i]
		if _, ok := added[aVal]; !ok {
			added[aVal] = true
			result = append(result, aVal)
		}
		if _, ok := added[bVal]; !ok {
			added[bVal] = true
			result = append(result, bVal)
		}
	}
	return result
}

// fmt.Println(uniqueNames(
// 	[]string{"Ava", "Emma", "Olivia"},
// 	[]string{"Olivia", "Sophia", "Emma"}))
