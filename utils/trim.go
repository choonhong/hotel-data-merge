package utils

import "strings"

func TrimSpaceInList(list []string) []string {
	var result []string
	for _, s := range list {
		result = append(result, strings.TrimSpace(strings.ToLower(s)))
	}
	return result
}
