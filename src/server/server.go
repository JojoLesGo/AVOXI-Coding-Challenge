package server

import (
	"net/http"

	"./controller"
)

// Server holds the information for the current status
type Server struct {
	baseURL     string
	middlewares []func(w http.ResponseWriter, r *http.Request)
}

// New creates the new server from the baseURL
func New(baseURL string) *Server {
	return &Server{
		baseURL: baseURL,
	}
}

// Middleware sets up the middleware for request processing
func (s *Server) Middleware(middleware func(w http.ResponseWriter, r *http.Request)) {
	s.middlewares = append(s.middlewares, middleware)
}

// Route configures the path for requests
func (s *Server) Route(path string, controller controller.Controller) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		for _, middleware := range s.middlewares {
			middleware(w, r)
		}

		id := r.URL.Path[1:]
		hasID := len(id) > 0

		switch r.Method {
		case "POST":
			if hasID {
				controller.PostOne(w, r, id)
			} else {
				controller.PostAll(w, r)
			}
		case "OPTIONS":
			controller.Options(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

// Serve starts the actual serving job
func (s *Server) Serve(port string) {
	http.ListenAndServe(":"+port, nil)
}
