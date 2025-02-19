package main

import (
	"LoadBalancer/backend"
	"LoadBalancer/healthcheck"
	"LoadBalancer/loadbalancer"
	"LoadBalancer/serverpool"
	"log"
)

func main() {
	backends := []*backend.Backend{
		backend.CreateBackend("http://localhost:6969", 1),
		backend.CreateBackend("http://localhost:4200", 2),
		backend.CreateBackend("http://localhost:3030", 3),
	}

	serverPool := serverpool.CreateServerPool(backends)

	healthcheck.StartHealthChecks(serverPool)

	l := loadbalancer.NewLoadBalancer(":8080", serverPool)
	log.Fatal(l.Start())
}
