package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	// setting the env variable for bot token and channel
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5114609547669-5117653705026-aS2ky99vU21Ql9yrW6LNJVTm")
	os.Setenv("CHANNEL_ID", "C053U58GVND")

	// creating new connection with bot
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// taking channels info into local variable
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	// saving file into local variable
	fileArr := []string{"sample.txt"}

	// iterating over all the files
	for i := 0; i < len(fileArr); i++ {

		// creating a slack file parameters varible to be used to call UploadFile() function
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		// uploading the actual file and getting slack file with information about the upload
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		// printing the uploaded file information
		fmt.Printf("Name: %s, URL:%s", file.Name, file.URL)
	}

}
