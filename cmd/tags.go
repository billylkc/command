package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var gitTagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"gt"},
	Short:   "[gt] Git tag next minor version.",
	Long:    `[gt] Git tag next minor version.`,
	Example: `
  command utility tag
  command u gt
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err, out, _ := myutil.Shell("git status")
		if err != nil {
			return err
		}

		// Check status first
		b := strings.Contains(out, "nothing to commit, working tree clean")
		if !b { // Need commit
			fmt.Println("Need to commit first. \n")
			fmt.Println("Exit")
			os.Exit(0)
		}

		// Get current tags
		err, out, _ = myutil.Shell(`git for-each-ref refs/tags --sort=-taggerdate --format='%(refname)' --count=1`)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
		// res := strings.Split(out, "\n")
		// current := res[len(res)-2] // current version
		// fmt.Printf("Current version: %s\n", current)

		// New tags
		// version, err := incrementVersion(current)
		// if err != nil {
		// 	return err
		// }

		// // Tag and push to origin
		// err, _, _ = myutil.Shell(fmt.Sprintf("git tag %s", version))
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("New version: %s\n", version)

		// err, out, _ = myutil.Shell(fmt.Sprintf("git push origin %s", version))
		// if err != nil {
		// 	return err
		// }

		// fmt.Printf("Push to origin.")
		// fmt.Println("")
		// fmt.Println(out)

		return nil
	},
}

func init() {
	utilityCmd.AddCommand(gitTagCmd)
}

func incrementVersion(s string) (string, error) {
	var version string

	re := "(v\\d\\.\\d\\.)(\\d)"
	regex := regexp.MustCompile(re)
	res := regex.FindStringSubmatch(s)

	if len(res) == 3 {
		_ = res[0]      // original tag, v0.0.4
		major := res[1] // first part, v0.0.
		minor := res[2] // second part 4
		minorVersion, err := strconv.Atoi(minor)
		if err != nil {
			return "", err
		}

		minorVersion += 1
		version = major + strconv.Itoa(minorVersion)
	} else {
		return "", fmt.Errorf("Invalid tag format. Should be something like this, v0.0.5.\nGot this, %s.\n", s)
	}
	return version, nil
}
