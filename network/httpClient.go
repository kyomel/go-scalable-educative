package network

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

var DefaultTimeout = 5

type httpClient struct {
	headers http.Header
	body    io.Reader
	timeout int
	name    string
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

func (client *httpClient) Do(method, url string) (*http.Response, error) {
	clientInstance := clientMap.getClient(client.name)
	if clientInstance == nil {
		clientInstance = clientMap.instantiateClient(client.name, client.timeout)
	}

	req, requestInitError := http.NewRequest(method, url, client.body)
	if requestInitError != nil {
		return nil, requestInitError
	}

	for k, v := range client.headers {
		if strings.ToLower(k) != "connection" {
			req.Header[k] = v
		}
	}

	res, err := clientInstance.Do(req)
	if err != nil || (res != nil && (res.StatusCode < 200 || res.StatusCode > 299)) {
		return res, errors.New("something went wrong")
	}

	return res, nil
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
