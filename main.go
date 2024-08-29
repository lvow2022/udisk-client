package main

import (
	udisk "udisk-client/udisk-sdk-go/services/udisk/v1"
	"udisk-client/udisk-sdk-go/test"
)

func main() {

	testUpload()
}

func testUpload() {
	client, _ := udisk.NewClient()
	client.SetBaseURL("http://localhost:8080")

	req := client.NewValidateUploadRequest()

	src := "/Users/luowei/Downloads/iTerm2-3_5_3.zip"
	dst := "/Users/luowei/go/src/udisk/download"
	req.SetSrc(src) // 本地
	req.SetDst(dst)

	resp, err := req.Send()
	if err != nil {
		return
	}

	chunkSize := resp.ChunkSize
	test.Upload(src, dst, chunkSize)
}
