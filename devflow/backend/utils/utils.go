package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatTag(tag string) string {
	tag = strings.ReplaceAll(strings.TrimSpace(tag), ".", " ")
	tag = cases.Title(language.English).String(tag)
	return tag
}