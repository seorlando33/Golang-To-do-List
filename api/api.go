package api

import (
	"net/http"
)


func New(addr string) *http.Server {
	
	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
