package command

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed templates/*
var templates embed.FS

// ParseMermaid parses the mermaid syntax to different Flow Chart diagram in a html files
// https://mermaid-js.github.io/mermaid/
func ParseMermaid(b []byte, w io.Writer) error {

	mermaid := Mermaid{
		Content: string(b),
	}
	t, err := template.ParseFS(templates, "templates/mermaid.gohtml")
	if err != nil {
		return err
	}

	err = t.Execute(w, mermaid)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(w)
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
