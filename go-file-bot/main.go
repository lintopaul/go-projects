package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	if os.Getenv("SLACK_BOT_TOKEN") == "" {
		fmt.Println("Env SLACK_BOT_TOKEN needs to be set")

		os.Exit(1)
	}
	
	client := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr := []string{"files/testfile", "files/testfile.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}

		file, err := client.UploadFile(params)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
