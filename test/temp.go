package test

import (
	"io/ioutil"
	"log"
	"os"
)

func InputValueCheck(input string, f func()) {
	//Tempファイルの作成
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	//Tempファイルへ書き込む
	content := []byte(input)
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	//Tempファイル情報をtmpfileに格納する
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
}
