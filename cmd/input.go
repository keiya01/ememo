package cmd

import (
	"errors"
	"strings"
)

func CheckingUserInputValue(args []string) error {
	if len(args) < 3 {
		return errors.New("ERROR: 引数を入力せずに実行することは出来ません。")
	}
	return nil
}

func AddExtension(fileName string) string {
	setFile := fileName
	isTxt := strings.HasSuffix(setFile, ".txt")
	if !isTxt {
		setFile += ".txt"
	}

	return setFile
}
