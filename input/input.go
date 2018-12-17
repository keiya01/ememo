package input

import (
	"bufio"
	"errors"
	"io/ioutil"
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

func InputValueCheck(input string, f func()) {
	//Tempファイルの作成
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	//Tempファイルへ書き込む
	content := []byte(input)
	if _, err := tmpfile.Write(content); err != nil {
		panic(err)
	}

	//Tempファイル情報をtmpfileに格納する
	if _, err := tmpfile.Seek(0, 0); err != nil {
		panic(err)
	}

	//もともとのos.Stdin情報をoldStdinに格納しておき、
	//最後にoldStdinをos.Stdinに代入して初期化する
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	//tmpfile情報をos.Stdinに代入することで、os.Stdinはポインタ型なので
	//関数内でos.Stdinを読み込むことで参照することができる
	os.Stdin = tmpfile

	f()

	if err := tmpfile.Close(); err != nil {
		panic(err)
	}
}
