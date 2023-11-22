package parse

import (
	"bufio"
	"fmt"
	"os"
)

func ParseSMS(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		fmt.Print(scanner.Text() + "\n")
	}
	return nil
}
