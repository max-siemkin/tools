package tools

import (
	"fmt"
	"strings"
)

type Tag struct {
	Tag     string
	Content any
	Attrs   map[string]string
}

func TagToStr(tags ...Tag) string {
	data := []string{}
	for _, tag := range tags {
		attr := ""
		fTag := tag.Tag
		for key, val := range tag.Attrs {
			atr := ""
			if val == "" {
				atr = key
			} else {
				atr = fmt.Sprintf(`%s="%s"`, key, val)
			}
			if attr == "" {
				attr = atr
			} else {
				attr = fmt.Sprintf("%s %s", attr, atr)
			}
		}
		if attr != "" {
			fTag = fmt.Sprintf("%s %s", tag.Tag, attr)
		}
		if tag.Content == nil {
			data = append(data, fmt.Sprintf("<%s/>", fTag))
		} else if c, ok := tag.Content.(Tag); ok {
			data = append(data, fmt.Sprintf("<%s>%v</%s>", fTag, TagToStr(c), tag.Tag))
		} else if c, ok := tag.Content.([]Tag); ok {
			data = append(data, fmt.Sprintf("<%s>%v</%s>", fTag, TagToStr(c...), tag.Tag))
		} else {
			data = append(data, fmt.Sprintf("<%s>%v</%s>", fTag, tag.Content, tag.Tag))
		}
	}
	return strings.Join(data, "")
}
