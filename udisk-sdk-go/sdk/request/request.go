package request

import (
	"github.com/go-resty/resty/v2"
	"udisk-client/udisk-sdk-go/sdk"
)

// Request is the base struct of service requests
type BaseRequest struct {
	*resty.Request
}

func NewBaseRequest(client *sdk.Client) *BaseRequest {
	return &BaseRequest{
		Request: client.NewRequest(),
	}
}
