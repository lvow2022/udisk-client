package v1

import (
	"encoding/json"
	"net/url"
	"udisk-client/udisk-sdk-go/sdk/request"
	"udisk-client/udisk-sdk-go/sdk/response"
)

// *****定义请求体和响应体的结构
type ValidateUploadResponse struct {
	*response.BaseResponse
	ChunkSize int64 `json:"chunk_size"`
}

type ValidateUploadRequest struct {
	*request.BaseRequest
}

func NewValidateUploadResponse() *ValidateUploadResponse {
	return &ValidateUploadResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

func (c *Client) NewValidateUploadRequest() *ValidateUploadRequest {
	r := request.NewBaseRequest(c.Client)
	return &ValidateUploadRequest{
		BaseRequest: r,
	}
}

func (r *ValidateUploadRequest) SetDst(dst string) {
	r.SetQueryParam("dst", url.QueryEscape(dst))
}

func (r *ValidateUploadRequest) SetSrc(src string) {
	r.SetQueryParam("src", url.QueryEscape(src))
}

func (r *ValidateUploadRequest) Send() (Resp *ValidateUploadResponse, err error) {
	rresp, err := r.Post("/file/validate/upload")
	resp := NewValidateUploadResponse()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rresp.Body(), resp)
	return resp, err
}
