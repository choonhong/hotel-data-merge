package utils

import (
	"strings"

	"github.com/agext/levenshtein"
)

// CalculateAverage calculates the average of a list of float64 numbers.
func CalculateAverage(list []float64) *float64 {
	if len(list) == 0 {
		return nil
	}

	total := 0.0
	for _, num := range list {
		total += num
	}

	average := total / float64(len(list))
	return &average
}

// FindMostRepresentativeString finds the most representative string from a list of strings.
func FindMostRepresentativeString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minTotalDistance := len(strs[0]) * len(strs) * 100
	mostAverageString := ""

	for _, str1 := range strs {
		totalDistance := 0

		for _, str2 := range strs {
			if str1 != str2 {
				totalDistance += levenshtein.Distance(str1, str2, nil)
			}
		}

		// all strings are the same
		if totalDistance == 0 {
			return str1
		}

		if totalDistance < minTotalDistance || (totalDistance == minTotalDistance && len(str1) > len(mostAverageString)) {
			minTotalDistance = totalDistance
			mostAverageString = str1
		}
	}

	return mostAverageString
}

// DedupStrings removes similar strings from a list of strings.
func DedupStrings(strs []string) []string {
	uniqueStrs := []string{}

	for _, str1 := range strs {
		unique := true

		for i := 0; i < len(uniqueStrs); i++ {
			str2 := uniqueStrs[i]
			if levenshtein.Distance(str1, str2, nil) < 2 || strings.Contains(str2, str1) || strings.Contains(str1, str2) {
				// both strings are similar or substring of the another, keep the longer one
				if len(str1) > len(str2) {
					uniqueStrs[i] = str1
				}

				unique = false
				break
			}
		}

		if unique {
			uniqueStrs = append(uniqueStrs, str1)
		}
	}

	return uniqueStrs
}
