package util

import (
	"bytes"
	"os/exec"
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

// Shell calls the shell command
func Shell(command string) (error, string, string) {
	const ShellToUse = "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
