package fakery

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type fakeryServer struct {
	port   string
	config *fakeryServerConfig
}

func CreateNewServer(port string, config *fakeryServerConfig) *fakeryServer {
	return &fakeryServer{
		port,
		config,
	}
}

func (s *fakeryServer) Start() {
	router := http.NewServeMux()

	for _, e := range s.config.Endpoints {
		ConfigureEndpoint(router, e)
	}

	log.Printf("Starting server on port %s\n", s.port)
	http.ListenAndServe(fmt.Sprintf(":%s", s.port), router)
}

func ConfigureEndpoint(router *http.ServeMux, endpoint FakeryEndpoint) {
	log.Printf("Creating endpoint %s %s\n", endpoint.Request.Method, endpoint.Request.Url)

	pattern := fmt.Sprintf("%s %s", endpoint.Request.Method, endpoint.Request.Url)
	router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(endpoint.Response.Status)

		for k, v := range endpoint.Response.Headers {
			w.Header().Set(k, v)
		}

		if endpoint.Response.Latency != 0 {
			time.Sleep(time.Duration(endpoint.Response.Latency) * time.Millisecond)
		}

		if endpoint.Response.Body != "" {
			w.Write([]byte(endpoint.Response.Body))
		}
	})
}
