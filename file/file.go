package file

import (
	"bufio"
	"os"
)

func PrintReadFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()

	return FileScan(file, func(scanner *bufio.Scanner, index int) string {
		return scanner.Text() + "\n"
	})
}

func FileScan(file *os.File, f func(scanner *bufio.Scanner, index int) string) string {
	var contents string
	index := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents += f(scanner, index)
		index++
	}

	return contents
}
