package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func ValidateDownload() {
	//发送 src，

	//接受 src 的 md5 和 chunksize
	//在 dst 目录接收文件
	//合并文件
}

// dw   form src  to dst
func Download(src string, dst string) {
	const chunkSize = 10 * 1024 * 1024
	md5 := "12312321b3jh12gf321td312hd3j12h"
	path := filepath.Join("/Users/luowei/go/src/udisk/download/", md5)
	os.Mkdir(path, 0750)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // 增加计数器
		go func(i int) {
			defer wg.Done() // 下载完成后减小计数器
			downloadChunk(path, i)
		}(i)
	}

	// 等待所有下载任务完成
	wg.Wait()
	fmt.Println("所有分片下载完成")
	merge(path, "/Users/luowei/go/src/udisk/download/item2.zip")
}

// dw /lw/example.txt /Users/luowei/go/src/udisk/download
// 表示上传的是二进制数据
func downloadChunk(path string, chunkIndex int) {
	url := "http://localhost:8080/file/download"
	chunkPath := filepath.Join(path, strconv.Itoa(chunkIndex))
	client := resty.New()

	resp, err := client.R().
		SetHeader("File-Path", "/lw/example.txt").
		SetHeader("Chunk-Index", strconv.Itoa(chunkIndex)).
		SetOutput(chunkPath). // 直接将字节数组作为请求体
		Get(url)              // 发送 POST 请求

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Failed to download chunk %d: %s", chunkIndex, resp.Status())
	}

	fmt.Printf("Successfully downloaded chunk %d\n", chunkIndex)
}

func merge(dir string, outputFile string) error {
	// 创建目标文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 需要合并的文件名列表
	files := []string{"0", "1", "2"}

	for _, fileName := range files {
		// 获取文件路径
		filePath := filepath.Join(dir, fileName)

		// 打开文件
		inFile, err := os.Open(filePath)
		if err != nil {
			return err
		}

		// 将文件内容拷贝到目标文件
		_, err = io.Copy(outFile, inFile)
		if err != nil {
			inFile.Close()
			return err
		}

		// 关闭当前文件
		inFile.Close()
	}

	return nil
}
