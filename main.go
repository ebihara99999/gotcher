package main

import (
	"fmt"
)

func main() {
	uploader := Uploader{Filename: "example.txt"}
	file, err := uploader.UploadToSlack()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
}
