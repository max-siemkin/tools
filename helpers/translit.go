package tools

import (
	"regexp"
	"strings"

	"github.com/rainycape/unidecode"
	"golang.org/x/text/language"
)

var reTranslit = regexp.MustCompile("[^a-zA-Z0-9]")
var replacer = map[language.Tag]map[string]string{
	language.AmericanEnglish: {
		"™": " Trademark ",
		"®": " Trademark ",
		"©": " Copyright ",
		"₽": " Ruble ",
		"$": " Dollar ",
		"€": " Euro ",
		"£": " Pound ",
		"₤": " Lira ",
		"¥": " Yen ",
		"&": " and ",
		"%": " Percent ",
		"№": " Numero ",
		"#": " Number ",
		"@": " At ",
		"°": " Degree ",
	},
	language.Russian: {
		"ь": "",
		"Ь": "",
		"ъ": "",
		"Ъ": "",
		"™": " TM ",
		"®": " TM ",
		"©": " Copyright ",
		"₽": " RUB ",
		"$": " USD ",
		"€": " EUR ",
		"£": " Pound ",
		"₤": " Lira ",
		"¥": " Yen ",
		"&": " i ",
		"%": " Procent ",
		"№": " Nomer ",
		"#": " Nomer ",
		"@": " Sobaka ",
		"°": " Gradus ",
	},
}

func Translit(tag language.Tag, lowercase bool, str string) string {
	str = unidecode.Unidecode(str)
	str = strings.TrimSpace(strings.Trim(str, ".,"))
	if r, ok := replacer[tag]; ok {
		for key, val := range r {
			str = strings.ReplaceAll(str, key, val)
		}
	}
	str = reTranslit.ReplaceAllString(str, " ")
	if lowercase {
		str = strings.ToLower(str)
	}
	return strings.Join(strings.Fields(str), "-")
}
