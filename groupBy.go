package statmachine

func GroupBy(res []Result, f func(Result) string) map[string][]Result {
	m := make(map[string][]Result)
	for _, r := range res {
		key := f(r)
		m[key] = append(m[key], r)
	}
	return m
}
