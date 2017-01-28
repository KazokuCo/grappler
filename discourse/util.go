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

func Excerpt(str string) string {
	parts := strings.SplitN(str, "\n\n", 2)
	switch len(parts) {
	case 0:
		return str
	case 1:
		return parts[0]
	default:
		return strings.TrimSuffix(parts[0], ".") + "…"
	}
}
