package serverpool

import (
	"LoadBalancer/backend"
	"math/rand"
)

type ServerPool struct {
	backends []*backend.Backend
}

// Create a new ServerPool
func CreateServerPool(backends []*backend.Backend) *ServerPool {
	return &ServerPool{
		backends: backends,
	}
}

// Getter for backends
func (s *ServerPool) GetBackends() []*backend.Backend {
	return s.backends
}

func (s *ServerPool) GetNextIndex() int {
	totalWeight := 0
	for _, b := range s.backends {
		if b.IsAlive() {
			totalWeight += b.Weight
		}
	}

	if totalWeight == 0 {
		return -1
	}

	randWeight := rand.Intn(totalWeight)
	cumulativeWeight := 0
	for i, b := range s.backends {
		if b.IsAlive() {
			cumulativeWeight += b.Weight
			if randWeight < cumulativeWeight {
				return i
			}
		}
	}
	return -1
}

func (s *ServerPool) GetNextPeer() *backend.Backend {
	nextIndex := s.GetNextIndex()
	if nextIndex == -1 {
		return nil
	}
	return s.backends[nextIndex]
}
