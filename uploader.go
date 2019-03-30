package main

import (
	"os"

	"github.com/nlopes/slack"
)

type Uploadable interface {
	UploadToSlack(filename string) (*slack.File, error)
}

type Uploader struct {
	Filename string
}

func (u *Uploader) UploadToSlack() (*slack.File, error) {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	currentDir, _ := os.Getwd()

	params := slack.FileUploadParameters{
		Title:    "Batman Example",
		File:     currentDir + "/" + u.Filename,
		Filename: u.Filename,
		Channels: []string{"home_jikka"},
	}
	file, err := api.UploadFile(params)
	return file, err
}
