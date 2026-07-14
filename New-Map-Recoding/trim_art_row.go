package main

import "strings"

func TrimArtRows(rows []string) []string {
	result := make([]string, len(rows))
	
	for k, i := range rows {
		result[k] = strings.TrimRight(i, " ")

	}
	return result
}
