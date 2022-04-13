package command

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/BurntSushi/toml"
)

func GenerateFurnitureExample(w io.Writer, impOnly bool) error {
	s := `
[[items]]
type = "book desk"
title= "电脑桌台式办公桌家用学生学习写字桌简易书桌简约卧室小户型桌子"
price=20.0
image="https://img.alicdn.com/imgextra/i2/2417352013/O1CN01LVEBWt1Qk01fugwRy_!!0-item_pic.jpg_430x430q90.jpg"
url="https://detail.tmall.com/item.htm?spm=a230r.1.14.7.327e4f8dHMhWlJ&id=571102285943&ns=1&abbucket=7"
important=true

[[items]]
type = "book shelf"
title= "电脑桌台式办公桌家用学生学习写字桌简易书桌简约卧室小户型桌子"
price=20.0
image="https://img.alicdn.com/imgextra/i2/2417352013/O1CN01LVEBWt1Qk01fugwRy_!!0-item_pic.jpg_430x430q90.jpg"
url="https://detail.tmall.com/item.htm?spm=a230r.1.14.7.327e4f8dHMhWlJ&id=571102285943&ns=1&abbucket=7"
important=false
`
	// Test if the toml file can be parsed to struct
	var (
		furniture Furniture
	)

	if _, err := toml.Decode(s, &furniture); err != nil {
		log.Fatal(err)
	}

	// Handle important flag
	if impOnly {
		var items []Item
		for _, item := range furniture.Items {
			if item.Important { // Important only
				items = append(items, item)
			}
			furniture = Furniture{
				Items: items,
			}
		}
		var buf bytes.Buffer
		err := toml.NewEncoder(&buf).Encode(furniture)
		if err != nil {
			return err
		}
		s = buf.String()

	} else {
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
