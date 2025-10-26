package t5

func IntersectInts(a, b []int) (bool, []int) {
	set := make(map[int]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}
	var result []int
	seen := make(map[int]struct{})
	for _, v := range b {
		if _, found := set[v]; found {
			if _, already := seen[v]; !already {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}
	return len(result) > 0, result
}
