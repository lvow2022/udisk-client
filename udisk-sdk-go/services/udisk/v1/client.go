package v1

import "udisk-client/udisk-sdk-go/sdk"

const (
	defaultEndpoint = "127.0.0.1:9090"
	serviceName     = "iam.authz"
)

type Client struct {
	*sdk.Client
}

func NewClient() (client *Client, err error) {
	client = &Client{
		Client: sdk.NewClient(),
	}
	client.Init(serviceName)
	return
}

func NewClientWithSecret(secretID, secretKey string) (client *Client, err error) {
	client = &Client{}
	config := sdk.NewConfig().WithEndpoint(defaultEndpoint)
	client.Init(serviceName).WithSecret(secretID, secretKey).WithConfig(config)
	return
}
