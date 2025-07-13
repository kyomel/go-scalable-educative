package network

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/kyomel/go-scalable-educative/logger"
	"go.uber.org/zap"
)

var DefaultTimeout = 1

type httpClient struct {
	headers http.Header
	body    io.Reader
	timeout int
	name    string
	ctx     *context.Context
}

func NewClient() *httpClient {
	return &httpClient{
		name:    "default",
		timeout: DefaultTimeout,
	}
}

func (client *httpClient) Headers(headers http.Header) *httpClient {
	if len(client.headers) == 0 {
		client.headers = make(map[string][]string)
	}
	for k, v := range headers {
		client.headers[k] = v
	}
	return client
}

func (client *httpClient) Body(body io.Reader) *httpClient {
	client.body = body
	return client
}

func (client *httpClient) Timeout(timeout int) *httpClient {
	client.timeout = timeout
	return client
}

func (client *httpClient) Name(name string) *httpClient {
	client.name = name
	return client
}

func (client *httpClient) WithContext(ctx *context.Context) *httpClient {
	client.ctx = ctx
	return client
}

func (client *httpClient) Do(method, url string) (*http.Response, error) {
	clientInstance := clientMap.getClient(client.name)
	if clientInstance == nil {
		clientInstance = clientMap.instantiateClient(client.name, client.timeout)
	}

	req, requestInitError := http.NewRequest(method, url, client.body)
	if requestInitError != nil {
		return nil, requestInitError
	}

	if client.ctx != nil {
		req = req.WithContext(*client.ctx)
	}

	for k, v := range client.headers {
		if strings.ToLower(k) != "connection" {
			req.Header[k] = v
		}
	}

	var res *http.Response
	var apiErr error

	err := hystrix.Do(client.name, func() error {
		res, apiErr = clientInstance.Do(req)
		if apiErr != nil {
			return apiErr
		}
		if res != nil && (res.StatusCode < 200 || res.StatusCode > 299) {
			return fmt.Errorf("non success status code found - %d", res.StatusCode)
		}
		return nil
	}, func(err error) error {
		logger.GetLoggerInstance().Error("API error", zap.Error(err))
		return fmt.Errorf("something went wrong")
	})

	return res, err
}

func (client *httpClient) Get(url string) (*http.Response, error) {
	return client.Do("GET", url)
}

func (client *httpClient) Put(url string) (*http.Response, error) {
	return client.Do("PUT", url)
}

func (client *httpClient) Post(url string) (*http.Response, error) {
	return client.Do("POST", url)
}
