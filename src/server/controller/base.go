package controller

import (
	"net/http"
)

// Controller is an abstract base type for MVC controllers
type Controller interface {
	PostOne(w http.ResponseWriter, r *http.Request, id string)
	PostAll(w http.ResponseWriter, r *http.Request)
	Options(w http.ResponseWriter, r *http.Request)
}
