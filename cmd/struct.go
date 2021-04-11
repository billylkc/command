package cmd

import "github.com/billylkc/myutil"

// Article as a general struct to print formatted topics with content
// usually as some colorized content with title, link, snippet combined in a single field
type Article struct {
	Date    string
	Content string
}
type Articles []Article

func (articles Articles) PrintTable() error {
	headers := []string{"Date", "Content"}
	ignores := []string{""}
	data := myutil.InterfaceSlice(articles)
	err := myutil.PrintTable(data, headers, ignores, 1)
	if err != nil {
		return err
	}
	return nil
}
