package network

import (
	"net/http"
	"sync"
	"time"
)

type clientManager struct {
	sync.RWMutex
	pool map[string]*http.Client
}

var clientMap clientManager
var transport http.Transport
var once sync.Once

func init() {
	once.Do(func() {
		defaultTransport, _ := http.DefaultTransport.(*http.Transport)
		transport = *defaultTransport
		transport.MaxIdleConns = 100
		transport.MaxIdleConnsPerHost = 10
		clientMap = clientManager{
			pool: make(map[string]*http.Client),
		}
	})
}

func (c *clientManager) getClient(key string) *http.Client {
	c.RLock()
	defer c.RUnlock()
	return c.pool[key]
}

func (c *clientManager) instantiateClient(key string, timeout int) *http.Client {
	c.Lock()
	defer c.Unlock()
	if c.pool[key] != nil {
		return c.pool[key]
	}
	c.pool[key] = &http.Client{
		Transport: &transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}
	return c.pool[key]
}
