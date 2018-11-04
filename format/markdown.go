package format

import (
	"fmt"
	"strings"
)

func ChengeToMarkdown(text string) string {
	var contents string
	sentences := strings.Split(text, ";")
	totalSentences := len(sentences)
	for _, sentence := range sentences {

		if sentence == "" {
			continue
		}

		switch string([]rune(sentence)[0]) {
		case "-":
			contents += strings.Replace(sentence, "-", " ● ", 1)
		case "=":
			contents += strings.Replace(sentence, "=", " ◎ ", 1)
		default:
			contents += sentence
		}

		contents += " [ ]"

		if totalSentences != 1 {
			contents += "\n"
		}

	}

	return contents
}

func ShowMarkdown() {
	fmt.Print("\n====== Markdown List ======\n\n")
	fmt.Print(" - : [ ● ] に変換されるので簡単にリストを作成できます\n")
	fmt.Print(" = : [ ◎ ] に変換されるので重要な項目に利用してください\n")
	fmt.Print(" ; : [ 改行 ] として認識されるので行の終わりに利用してください\n")
	fmt.Print("\n====== END ======\n")
}
