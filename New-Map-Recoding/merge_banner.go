package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)

	for k, v := range base {
		result[k] = v
	}

	for k, v := range priority {
		result[k] = v
	}
	return result
}
