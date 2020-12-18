package main

import (
	"net/http"
	"os"

	"../server"
	"../server/controller"
	"../server/middleware"

	model "../server/model"
)

var (
	port    = "8080"
	baseURL = "http://localhost:" + port
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("BASE_URL"); env != "" {
		baseURL = env
	}
}

func main() {
	i := new(model.Info)

	c := controller.Info(*i)
	s := server.New(baseURL)

	s.Middleware(func(w http.ResponseWriter, r *http.Request) { middleware.Cors(w) })

	s.Route("/", c)

	s.Serve(port)
}
