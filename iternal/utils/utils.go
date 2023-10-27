package utils

import "strings"

func IsStringContainsInSlice(sample string, slice []string) bool {
	for i := range slice {

		if strings.EqualFold(strings.ToLower(sample), strings.ToLower(slice[i])) {
			return true
		}
	}
	return false
}
