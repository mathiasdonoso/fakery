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
	log.Printf("Starting server at port %s\n", s.port)

	router := http.NewServeMux()

	for _, e := range s.config.Endpoints {
		ConfigureEndpoint(router, e)
	}

	log.Printf("Starting server on port %s\n", s.port)
	http.ListenAndServe(fmt.Sprintf(":%s", s.port), router)
}

func ConfigureEndpoint(router *http.ServeMux, endpoint FakeryEndpoint) {
	req := endpoint.Request
	res := endpoint.Response

	pattern := fmt.Sprintf("%s %s", req.Method, req.Url)
	router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		for k, v := range res.Headers {
			w.Header().Set(k, v)
		}

		w.WriteHeader(res.Status)

		if res.Latency != 0 {
			time.Sleep(time.Duration(res.Latency) * time.Millisecond)
		}

		if res.Body != "" {
			w.Write([]byte(res.Body))
		}
	})
}
