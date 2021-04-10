package util

import (
	"regexp"
	"strings"

	. "github.com/logrusorgru/aurora"
)

// ParseQuery simply joins the arguments with underscore
// Used to join wiki query strgin
func ParseQuery(q ...string) string {
	query := strings.Join(q, "_")
	return query
}

func InsensitiveReplace(s, search string, once bool) string {
	var output string
	replace := Sprintf(Red(search))
	// replace := "[" + search + "]"
	flag := false
	pat := regexp.MustCompile("(?i)(" + search + ")")
	if once {
		output = pat.ReplaceAllStringFunc(s, func(a string) string {
			if flag {
				return a
			}
			flag = true
			return pat.ReplaceAllString(a, replace)
		})
	} else {
		output = pat.ReplaceAllString(s, replace)
	}
	return output
}
