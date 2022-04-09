package command

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/BurntSushi/toml"
)

func GenerateFurnitureExample(w io.Writer) error {
	s := `
[[items]]
type = "book desk"
title= "电脑桌台式办公桌家用学生学习写字桌简易书桌简约卧室小户型桌子"
image = "https://img.alicdn.com/imgextra/i2/2417352013/O1CN01LVEBWt1Qk01fugwRy_!!0-item_pic.jpg_430x430q90.jpg"
url = "https://detail.tmall.com/item.htm?spm=a230r.1.14.7.327e4f8dHMhWlJ&id=571102285943&ns=1&abbucket=7"
`
	// Test if the toml file can be parsed to struct
	var furniture Furniture
	if _, err := toml.Decode(s, &furniture); err != nil {
		log.Fatal(err)
	}

	// Write content to writter
	writer := bufio.NewWriter(w)
	writer.WriteString(s)
	err := writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

	return nil

}
