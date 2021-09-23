package main

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

func initRegexPatterns(patterns map[string]string) (m map[string]*regexp.Regexp, r *regexp.Regexp) {
	var fullPattern string
	m = make(map[string]*regexp.Regexp)

	if len(patterns) > 0 {
		for key, val := range patterns {
			m[key] = regexp.MustCompile(val)
			fullPattern += fmt.Sprint("|", val)
		}
		r = regexp.MustCompile(fullPattern[1:])
	}

	return
}

func Translate(text string) (parsed string) {
	var (
		// Regexp
		rgx, rgxToken = initRegexPatterns(
			map[string]string{
				"hr":     `(^|\n) ?-{3} ?(\n|$)`,
				"h":      `(^|\n)#{1,7} ?`,
				"a":      `\[.+\]\([\w\./:%&\?!=-]+\)`,
				"format": `[_*]{1,2}`,
			},
		)
		rgxEndln = regexp.MustCompile(`\r?\n`)

		tag     string
		cursor  int
		wait    = make(map[string]bool)
		tagName = map[string]string{
			"_":  "i",
			"*":  "b",
			"__": "em",
			"**": "strong",
		}
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

		case rgx["h"].MatchString(token):
			tag = fmt.Sprint("<h", len(strings.TrimSpace(token)), ">")
			if val := rgxEndln.FindStringIndex(text[max:]); val != nil {
				parsed += html.EscapeString(text[cursor:min]) + tag
				parsed += Translate(text[max:max+val[0]]) + "</" + tag[1:] + "\n"
				max = max + val[1]
				cursor, min = max, max
				continue
			}

		case rgx["a"].MatchString(token):
			link := strings.Split(token[1:len(token)-1], "](")
			if match, _ := regexp.MatchString(`(?i).*\.(apng|avif|gif|jpe?g|jpe|jf?if|png|svg|webp)$`, link[1]); match {
				tag = fmt.Sprint("<img src=\"", link[1], "\" alt=\"", html.EscapeString(link[0]), "\">")
			} else {
				tag = fmt.Sprint("<a href=\"", link[1], "\">", Translate(link[0]), "</a>")
			}

		case rgx["hr"].MatchString(token):
			tag = "\n<hr>\n"

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
