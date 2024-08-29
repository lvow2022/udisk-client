package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

// ul from local src to remote dst
func Upload(src string, dst string, chunkSize int64) {
	fileMd5, err := calculateFileMD5(src)
	// 计算分片数量
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	totalSize := fileInfo.Size()
	numChunks := int(math.Ceil(float64(totalSize) / float64(chunkSize)))

	//每个分片数量开启一个 goroutine 发送
	for i := 0; i < numChunks; i++ {
		offset := int64(i) * chunkSize
		remaining := totalSize - offset
		currentChunkSize := chunkSize
		if remaining <= chunkSize {
			currentChunkSize = remaining
		}

		chunk := make([]byte, currentChunkSize)
		_, err := file.ReadAt(chunk, offset)
		if err != nil {
			fmt.Println("Error reading chunk:", err)
			return
		}

		ChunkMd5 := calculateMD5(chunk)
		// 上传当前分片
		uploadChunk(chunk, i, ChunkMd5, fileMd5)
	}

	// 发送 complete
	url := "http://localhost:8080/file/complete"
	client1 := resty.New()
	client1.R().
		SetQueryParam("file_md5", fileMd5).
		SetQueryParam("chunk_num", strconv.Itoa(numChunks)).
		Post(url)

}

// 表示上传的是二进制数据
func uploadChunk(chunk []byte, chunkIndex int, ChunkMd5, FileMd5 string) {
	url := "http://localhost:8080/file/upload"
	client := resty.New()
	resp, err := client.R().
		SetHeader("File-Md5", FileMd5).
		SetHeader("Content-Type", "application/octet-stream"). // 设置内容类型
		SetHeader("Chunk-Index", strconv.Itoa(chunkIndex)).
		SetHeader("Chunk-Md5", ChunkMd5).
		SetBody(chunk). // 直接将字节数组作为请求体
		Post(url)       // 发送 POST 请求

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Failed to upload chunk %d: %s", chunkIndex, resp.Status())
	}

	fmt.Printf("Successfully uploaded chunk %d\n", chunkIndex)
}

func calculateFileMD5(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建 MD5 哈希器
	hash := md5.New()

	// 将文件的内容写入哈希器
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// 计算哈希值
	hashInBytes := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

func calculateMD5(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
