package input

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/fatih/color"
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

func GetUserInputValue() (string, error) {
	// Scannerを使って一行読み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := scanner.Text()
	if inputText == "" {
		return "", errors.New("入力値を空にすることは出来ません")
	}

	color.HiGreen("input value is [ %s ]", inputText)
	return inputText, nil
}
