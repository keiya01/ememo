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

	if string([]rune(text)[0:2]) == "[]" {
		sentence = strings.Replace(text, "[]", "[ ] ", 1) + "\n"
	}

	sentence += "\n"

	return sentence
}

func ShowMarkdown() {
	fmt.Print("\n====== Markdown List ======\n\n")
	fmt.Print(" - : [ ● ] に変換されるので簡単にリストを作成できます\n")
	fmt.Print(" = : [ ◎ ] に変換されるので重要な項目に利用してください\n")
	fmt.Print(" [] : TODOリストを作成したいときに利用してください\n=> --compフラグを使うことで完了を明示的にあらわす事が出来ます\n")
	fmt.Print(" end : [ 入力終了 ] として認識されます\n")
	fmt.Print("\n====== END ======\n")
}
