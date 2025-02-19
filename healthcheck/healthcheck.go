package healthcheck

import (
	"LoadBalancer/serverpool"
	"log"
	"net/http"
	"net/url"
	"time"
)

func IsBackendAlive(url *url.URL) bool {
	resp, err := http.Get(url.String())
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func StartHealthChecks(serv *serverpool.ServerPool) {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			for _, b := range serv.GetBackends() {
				alive := IsBackendAlive(b.URL)
				b.SetAlive(alive)
				status := "up"
				if !alive {
					status = "down"
				}
				log.Printf("%s [%s]\n", b.URL, status)
			}
		}
	}()
}
