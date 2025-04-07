package services

import (
	"net/http"
	"sync"
	"time"
)

var (
	client     *http.Client
	clientOnce sync.Once
)

func getClient() *http.Client {
	clientOnce.Do(func() {
		client = &http.Client{
			Timeout: 15 * time.Second,
			// Transport можно настроить тут (прокси, TLS и т.д.)
		}
	})
	return client
}
