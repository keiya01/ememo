package format

import (
	"fmt"
	"strings"
)

func ChangeToMarkdown(text string, isNextLine bool) string {

	if text == "\n" || len(text) == 0 {
		return "\n"
	}

	var sentence string

	switch string([]rune(text)[0:2]) {
	case "[]":
		sentence = strings.Replace(text, "[]", "[ ] ", 1)
	case "--":
		sentence = strings.Replace(text, "--", "   ○ ", 1)
	}

	if sentence == "" {
		switch string([]rune(text)[0]) {
		case "-":
			sentence = strings.Replace(text, "-", " ● ", 1)
		case "=":
			sentence = strings.Replace(text, "=", " ◎ ", 1)
		default:
			sentence = text
		}
	}

	if isNextLine {
		sentence += "\n"
	}

	return sentence
}

func ShowMarkdown() {
	fmt.Print("\n====== Markdown List ======\n\n")
	fmt.Print(" - : [ ● ] に変換されるので簡単にリストを作成できます\n")
	fmt.Print(" -- : [ ○ ] に変換されるので簡単にリストを作成できます。この文字は`-`に続けて利用します。\n")
	fmt.Print(" = : [ ◎ ] に変換されるので重要な項目に利用してください\n")
	fmt.Print(" [] : TODOリストを作成したいときに利用してください。--compフラグを使うことで完了を明示的にあらわす事が出来ます。\n")
	fmt.Print(" end : [ 入力終了 ] として認識されます\n")
	fmt.Print(" -e line_number : line_numberに指定した行を編集します\n")
	fmt.Print(" -d line_number : line_numberに指定した行を削除します\n")
	fmt.Print(" -s : 入力したテキストにマークダウンが適用されたものを確認できます\n")
	fmt.Print("\n====== END ======\n")
}
