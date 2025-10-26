package t4

func Difference(slice1, slice2 []string) []string {
	set := make(map[string]struct{}, len(slice2))
	for _, v := range slice2 {
		set[v] = struct{}{}
	}
	var result []string
	for _, v := range slice1 {
		if _, found := set[v]; !found {
			result = append(result, v)
		}
	}
	return result
}
