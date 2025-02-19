package serverpool

import (
	"LoadBalancer/backend"
	"testing"
)

func TestWeightedRoundRobin(t *testing.T) {
	// Create test backends with weights
	backends := []*backend.Backend{
		backend.CreateBackend("http://localhost:6969", 1),
		backend.CreateBackend("http://localhost:4200", 2),
		backend.CreateBackend("http://localhost:3030", 3),
	}

	// Create a server pool
	serverPool := CreateServerPool(backends)

	// Test weighted selection
	count := make(map[string]int)
	totalRequests := 6000
	for i := 0; i < totalRequests; i++ {
		peer := serverPool.GetNextPeer()
		if peer == nil {
			t.Error("Expected a backend, got nil")
			continue
		}
		count[peer.URL.String()]++
	}

	// Verify distribution (approximate)
	expected := map[string]int{
		"http://localhost:6969": totalRequests * 1 / 6,
		"http://localhost:4200": totalRequests * 2 / 6,
		"http://localhost:3030": totalRequests * 3 / 6,
	}

	tolerance := 0.1
	for url, cnt := range count {
		expectedCnt := expected[url]
		lowerBound := int(float64(expectedCnt) * (1 - tolerance))
		upperBound := int(float64(expectedCnt) * (1 + tolerance))

		if cnt < lowerBound || cnt > upperBound {
			t.Errorf("Expected %s to handle ~%d requests (with tolerance), got %d", url, expectedCnt, cnt)
		}
	}
}
