package network

import (
	"net/http"
	"sync"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

type clientManager struct {
	sync.RWMutex
	pool map[string]*http.Client
}

const (
	DEFAULT_MAX_CONCURRENCY  = 10
	DEFAULT_ERROR_THRESHOLD  = 25
	DEFAULT_VOLUME_THRESHOLD = 5
	DEFAULT_SLEEP_WINDOW     = 3000
)

var (
	clientMap clientManager
	transport http.Transport
	once      sync.Once
)

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
	hystrix.ConfigureCommand(key, hystrix.CommandConfig{
		Timeout:                timeout * 1000,
		MaxConcurrentRequests:  DEFAULT_MAX_CONCURRENCY,
		RequestVolumeThreshold: DEFAULT_VOLUME_THRESHOLD,
		ErrorPercentThreshold:  DEFAULT_ERROR_THRESHOLD,
		SleepWindow:            DEFAULT_SLEEP_WINDOW,
	})
	return c.pool[key]
}
