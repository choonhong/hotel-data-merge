package utils

import "github.com/agext/levenshtein"

// CalculateAverage calculates the average of a list of float64 numbers.
func CalculateAverage(list []float64) float64 {
	total := 0.0
	for _, num := range list {
		total += num
	}
	return total / float64(len(list))
}

// FindMostAverageString combines multiple strings into one, returning the most average string
func FindMostAverageString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minTotalDistance := len(strs[0]) * len(strs) * 100
	averageString := ""

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

		if totalDistance < minTotalDistance || (totalDistance == minTotalDistance && len(str1) > len(averageString)) {
			minTotalDistance = totalDistance
			averageString = str1
		}
	}

	return averageString
}

// GetUniqueStrings return a list of unique strings where levenshtein distance is < 3
func GetUniqueStrings(strs []string) []string {
	uniqueStrs := []string{}

	for _, str1 := range strs {
		unique := true

		for _, str2 := range uniqueStrs {
			if levenshtein.Distance(str1, str2, nil) < 3 {
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
