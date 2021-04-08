package util

import (
	"regexp"
	"strings"
	// . "github.com/logrusorgru/aurora"
)

// ParseQuery simply joins the arguments with underscore
// Used to join wiki query strgin
func ParseQuery(q ...string) string {
	query := strings.Join(q, "_")
	return query
}

func InsensitiveReplace(s, search string) string {
	// replace := fmt.Sprintf("Cyan(%s)", search)
	replace := "[" + search + "]"
	searchRegex := regexp.MustCompile("(?i)" + search)
	return searchRegex.ReplaceAllString(s, replace)
}
