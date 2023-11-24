package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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
func CopyFileToOutputDirectory(src string, dst string) (int64, error) {

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	dst = filepath.Join(dst, sourceFileStat.Name())
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
