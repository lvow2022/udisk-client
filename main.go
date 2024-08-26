package main

import (
	"fmt"
	udisk "udisk-client/udisk-sdk-go/services/udisk/v1"
)

func main() {
	testValidateUpload()
}

func testValidateUpload() {
	client, _ := udisk.NewClient()
	client.SetBaseURL("http://localhost:8080")

	m := &udisk.Metadata{
		Name:    "example.txt",
		Size:    1024000000,
		Type:    "text/plain",
		OwnerID: "user_12345",
		MD5:     "d41d8cd98f00b204e9800998ecf8427e",
		Path:    "/example.txt",
	}
	req := client.NewValidateUploadRequest()
	req.FileMetadata(m)

	err := req.Send()
	if err != nil {
		fmt.Println("err1", err)
		return
	}
}
