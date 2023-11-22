package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func IsStringContainsInSlice(sample string, slice []string) bool {
	fmt.Println("slice:", slice)
	for i := range slice {
		fmt.Println("method:", strings.ToLower(sample), strings.ToLower(slice[i]))
		if strings.EqualFold(strings.ToLower(sample), strings.ToLower(slice[i])) {
			return true
		}
	}
	return false
}

func GetFileExtensionFromPath(path string) string {
	fmt.Println("ext Path:", path)
	fmt.Println("ext:", path[strings.LastIndex(path, ".")+1:])
	return path[strings.LastIndex(path, ".")+1:]
}

func IsMultipleFileEvents(path string) (error, bool) {
	current_time := time.Now()
	sourceFileStat, err := os.Stat(path)
	fmt.Println(sourceFileStat.ModTime())
	fmt.Println(current_time)
	if err != nil {
		return err, true
	}
	if sourceFileStat.ModTime().Compare(current_time) == 0 {
		return nil, true
	}

	return nil, false
}
