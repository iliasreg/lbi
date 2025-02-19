package loadbalancer

import (
	"LoadBalancer/serverpool"
	"net/http"
)

type LoadBalancer struct {
	port       string
	serverPool *serverpool.ServerPool
}

func NewLoadBalancer(port string, s *serverpool.ServerPool) *LoadBalancer {
	return &LoadBalancer{
		port:       port,
		serverPool: s,
	}
}

func (l *LoadBalancer) Start() error {
	http.HandleFunc("/", l.HandleRequest)
	return http.ListenAndServe(l.port, nil)
}

func (l *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	peer := l.serverPool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
