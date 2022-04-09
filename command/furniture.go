package command

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"text/template"

	"github.com/BurntSushi/toml"
)

func ParseFurniture(b []byte, w io.Writer) error {

	// Test if the toml file can be parsed to struct
	var furniture Furniture
	if _, err := toml.Decode(string(b), &furniture); err != nil {
		log.Fatal(err)
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
