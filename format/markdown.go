package format

import (
	"fmt"
	"strings"
)

func ChengeToMarkdown(text string) string {
	var sentence string

	switch string([]rune(text)[0]) {
	case "-":
		sentence = strings.Replace(text, "-", " ● ", 1)
	case "=":
		sentence = strings.Replace(text, "=", " ◎ ", 1)
	default:
		sentence = text
	}

	sentence += "[ ]\n"

	return sentence
}

func ShowMarkdown() {
	fmt.Print("\n====== Markdown List ======\n\n")
	fmt.Print(" - : [ ● ] に変換されるので簡単にリストを作成できます\n")
	fmt.Print(" = : [ ◎ ] に変換されるので重要な項目に利用してください\n")
	fmt.Print(" end : [ 入力終了 ] として認識されるので入力が終了したら`end`と入力してください\n")
	fmt.Print("\n====== END ======\n")
}
