package command

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
)

func ParseFurniture(b []byte, w io.Writer) error {

	// Test if the toml file can be parsed to struct
	var furniture Furniture
	if _, err := toml.Decode(string(b), &furniture); err != nil {
		log.Fatal(err)
	}

	// Fill title if empty
	for i, item := range furniture.Items {
		if strings.TrimSpace(item.Title) == "" {
			furniture.Items[i].Title = "Here"
		}
	}

	t, err := template.ParseFS(templates, "templates/furniture.gohtml")
	if err != nil {
		return err
	}

	err = t.Execute(w, furniture.Items)
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
