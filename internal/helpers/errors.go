package helpers

import (
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteErrorJSON(w, "the server encountered a problem", http.StatusInternalServerError)
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
}

func NotFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
}
