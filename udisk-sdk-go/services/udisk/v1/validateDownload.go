package v1

import (
	"fmt"
	"net/url"
	"udisk-client/udisk-sdk-go/sdk/request"
	"udisk-client/udisk-sdk-go/sdk/response"
)

// *****定义请求体和响应体的结构
type ValidateDownloadResponse struct {
	*response.BaseResponse
	Md5        string `json:"md_5"`
	ChunkCount int    `json:"chunk_count"`
}

type ValidateDownloadRequest struct {
	*request.BaseRequest
}

func NewValidateDownloadResponse() *ValidateDownloadResponse {
	return &ValidateDownloadResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

func (c *Client) NewValidateDownloadRequest() *ValidateDownloadRequest {
	r := request.NewBaseRequest(c.Client)
	return &ValidateDownloadRequest{
		BaseRequest: r,
	}
}

func (r *ValidateDownloadRequest) SetDst(dst string) {
	r.SetQueryParam("dst", url.QueryEscape(dst))
}

func (r *ValidateDownloadRequest) SetSrc(src string) {
	r.SetQueryParam("src", url.QueryEscape(src))
}

func (r *ValidateDownloadRequest) Send() error {
	resp, err := r.Post("/file/validate/download")
	if err != nil {
		return err
	}
	fmt.Printf("resp : %v\n", resp.String())
	return nil
}
