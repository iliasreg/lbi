package backend

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	Weight       int
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func CreateBackend(link string, weight int) *Backend {
	// Creating URL based on the provided input string
	u, err := url.Parse(link)
	if err != nil {
		fmt.Println(err)
	}
	// Initializing the Reverse Proxy for the created URL
	reverseProxy := httputil.NewSingleHostReverseProxy(u)

	return &Backend{
		URL:          u,
		Alive:        true,
		Weight:       weight,
		ReverseProxy: reverseProxy,
	}
}

// Setter for backend status
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// Getter for backend status
func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	alive := b.Alive
	b.mux.RUnlock()
	return alive
}
