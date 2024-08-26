package sdk

import (
	"github.com/go-resty/resty/v2"
	"time"
	"udisk-client/udisk-sdk-go/sdk/log"
)

type Request = resty.Request

// Client 封装了 resty.Client 并提供了一些常用方法
type Client struct {
	Config      *Config
	ServiceName string
	Logger      log.Logger
	client      *resty.Client
}

// NewClient 创建一个新的封装客户端实例
func NewClient() *Client {
	return &Client{
		client: resty.New(),
	}
}

func (c *Client) Init(serviceName string) *Client {
	c.Logger = log.New()
	c.ServiceName = serviceName
	return c
}

func (c *Client) WithCredential() *Client {

	return c
}

func (c *Client) WithSecret(secretID, secretKey string) *Client {
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	c.Logger.SetLevel(config.LogLevel)
	return c
}

// SetBaseURL 设置基础 URL
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.client.SetBaseURL(baseURL)
	return c
}

// SetTimeout 设置请求超时时间
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.client.SetTimeout(timeout)
	return c
}

// SetHeader 设置请求头
func (c *Client) SetHeader(key, value string) *Client {
	c.client.SetHeader(key, value)
	return c
}

// SetAuth 设置基本认证
func (c *Client) SetAuth(username, password string) *Client {
	c.client.SetBasicAuth(username, password)
	return c
}

// SetDebug 开启调试模式
func (c *Client) SetDebug(enable bool) *Client {
	c.client.SetDebug(enable)
	return c
}

// SetRetry 设置重试策略
func (c *Client) SetRetry(count int, wait time.Duration) *Client {
	c.client.SetRetryCount(count)
	c.client.SetRetryWaitTime(wait)
	return c
}

func (c *Client) NewRequest() *resty.Request {
	return c.client.R()
}
