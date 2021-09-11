package main

import (
	"html"
	"regexp"
	"strings"
)

func translate(text string) (parsed string) {
	var (
		tag     string
		cursor  int
		wait    = make(map[string]bool)
		tagName = map[string]string{
			"_":  "i",
			"*":  "b",
			"#":  "h1",
			"__": "em",
			"**": "strong",
			"##": "h2",
		}
		rgxToken = regexp.MustCompile(`[_*#]+`)
		rgxEndln = regexp.MustCompile(`\r?\n`)
	)

	vals := rgxToken.FindAllStringIndex(text, -1)
	for _, ind := range vals {
		min, max := ind[0], ind[1]
		token := text[min:max]

		if wait[token] {
			tag = "</" + tagName[token] + ">"
			delete(wait, token)
		} else {
			tag = "<" + tagName[token] + ">"
			wait[token] = true
		}

		parsed += html.EscapeString(text[cursor:min]) + tag
		cursor = max
		if token[0] == '#' {
			line := ""
			if val := rgxEndln.FindStringIndex(text[cursor:]); val != nil {
				line = text[cursor:val[1] + cursor]
				cursor += val[1]
			} else {
				line = text[cursor:]
				cursor += len(line)
			}
			parsed += parse(strings.TrimSpace(line)) + "</" + tag[1:] + "\n"
			delete(wait, token)
		}
	}
	parsed += html.EscapeString(text[cursor:])

	return
}
