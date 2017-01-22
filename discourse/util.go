package discourse

import (
	"regexp"
	"strconv"
	"strings"
)

var HTMLRE = regexp.MustCompile(`<([^>]+)>`)

func AvatarURL(tpl string, maxSize int) string {
	return strings.Replace(tpl, "{size}", strconv.Itoa(maxSize), -1)
}

func StripHTML(str string) string {
	return HTMLRE.ReplaceAllString(str, "")
}
