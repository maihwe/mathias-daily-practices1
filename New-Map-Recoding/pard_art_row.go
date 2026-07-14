package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	if width <= 0 {
		return rows
	}
	res := make([]string, len(rows))
	for i, v := range rows {
		padding := width - len(v)
		if padding > 0 {
			res[i] = v + strings.Repeat(" ", padding)
		} else {
			res[i] = v
		}
	}
	return res
}
