package lib

import (
	"fmt"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%v", err)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%v", err)
}

func BadGateway(w http.ResponseWriter){
	w.WriteHeader(http.StatusMethodNotAllowed)
}