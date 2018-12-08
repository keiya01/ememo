package input

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func CheckingUserInputArgumentValue(args []string) error {
	if len(args) < 2 {
		return errors.New("引数を2つ以上入力してください。")
	}
	return nil
}

func AddExtension(fileName string) string {
	spliteName := strings.Split(fileName, ".")
	setFile := spliteName[0] + ".txt"

	return setFile
}

func GetUserInputValue() string {
	// Scannerを使って一行読み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := scanner.Text()

	return inputText
}
