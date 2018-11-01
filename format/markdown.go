package format

import (
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

		if totalSentences != 1 {
			contents += "\n"
		}

	}

	return contents
}
