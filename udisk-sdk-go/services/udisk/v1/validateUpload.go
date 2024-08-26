package v1

import (
	"fmt"
	"udisk-client/udisk-sdk-go/sdk/request"
	"udisk-client/udisk-sdk-go/sdk/response"
)

type TaskType int

const (
	RegularFileUpload TaskType = iota
	LargeFileUpload
	RegularFileDownload
	LargeFileDownload
)

// 请求体
// todo  这部分应该由 udisk 导入
type Metadata struct {
	Name    string `json:"name"`     // 文件名称
	Size    int    `json:"size"`     // 文件大小 (字节)
	Type    string `json:"type"`     // 文件类型 (例如 "image/png")
	OwnerID string `json:"owner_id"` // 上传者ID (例如用户ID)
	MD5     string `json:"md5"`      // 文件内容的MD5哈希值
	Path    string `json:"path"`
}

// 响应体

type TaskDetails struct {
	TaskId   string      `json:"task_id"`
	Type     TaskType    `json:"type"`
	ChunkNum int         `json:"chunk_num"`
	Chunks   []ChunkInfo `json:"chunks"`
}
type ChunkInfo struct {
	Index int `json:"index"`
	Start int `json:"start"`
	End   int `json:"end"`
}

// *****定义请求体和响应体的结构
type ValidateUploadResponse struct {
	*response.BaseResponse
	Task_details TaskDetails `json:"task_details"`
}

type ValidateUploadRequestBody struct {
	Metadata *Metadata `json:"metadata"`
}

type ValidateUploadRequest struct {
	*request.BaseRequest
	body *ValidateUploadRequestBody
}

// *****定义请求体和响应体的构造方法
func NewValidateUploadResponse() *ValidateUploadResponse {
	return &ValidateUploadResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

func (c *Client) NewValidateUploadRequest() *ValidateUploadRequest {
	r := request.NewBaseRequest(c.Client)
	return &ValidateUploadRequest{
		BaseRequest: r,
		body:        &ValidateUploadRequestBody{},
	}
}

func (r *ValidateUploadRequest) FileMetadata(m *Metadata) {
	r.body.Metadata = m
	r.SetBody(r.body)
}

func (r *ValidateUploadRequest) Send() error {
	resp, err := r.Post("/file/validate/upload")
	if err != nil {
		return err
	}
	fmt.Printf("resp : %v\n", resp.String())
	return nil
}
