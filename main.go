package main

import (
	"html"
	"regexp"
)

func Translate(text string) (parsed string) {
	var (
		tag     string
		cursor  int
		wait    = make(map[string]bool)
		tagName = map[string]string{
			"_":  "i",
			"*":  "b",
			"__": "em",
			"**": "strong",
		}
		rgxToken = regexp.MustCompile(`[_*#]+`)
		rgxEndln = regexp.MustCompile(`\r?\n`)
	)

	vals := rgxToken.FindAllStringIndex(text, -1)
	for _, ind := range vals {
		if cursor > ind[0] {
			continue
		}
		min, max := ind[0], ind[1]
		token := text[min:max]

		switch true {
		case wait[token]:
			tag = "</" + tagName[token] + ">"
			delete(wait, token)

		case token[0] == '#':
			tag = fmt.Sprint("<h", len(token), ">")
			if val := rgxEndln.FindStringIndex(text[max:]); val != nil {
				parsed += html.EscapeString(text[cursor:min]) + tag
				parsed += parse(text[max:max + val[0]]) + "</" + tag[1:] + "\n"
				max = max + val[1]
				cursor, min = max, max
				continue
			}

		default:
			tag = "<" + tagName[token] + ">"
			wait[token] = true
		}

		parsed += html.EscapeString(text[cursor:min]) + tag
		cursor = max
	}
	parsed += html.EscapeString(text[cursor:])

	return
}
