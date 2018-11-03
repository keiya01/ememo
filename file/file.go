package file

import (
	"bufio"
	"fmt"
	"os"
)

func PrintReadFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	return fileScan(file)
}

func fileScan(file *os.File) string {
	var contents string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = contents + scanner.Text() + "\n"
	}

	return contents
}
