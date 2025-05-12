package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-simple-restfull/handle"
	"go-simple-restfull/middleware"
	"net/http"
)

func initServer() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("Root path is not allowed")
	})

	// Global Error Handler
	GlobalErrorHandler(router)

	// simple, just playing with routers
	handle.RegisterTestSimpleRoutes(router)
	handle.RegisterTestSimpleRoutesWithFS(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &middleware.LogMiddleware{router},
	}

	server.ListenAndServe()
}

func GlobalErrorHandler(router *httprouter.Router) {
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		ContentTypeJson(w)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"code": "500", "message": "%s"}`, err)
	}

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ContentTypeJson(w)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"code": "404", "message": "Not Found, please check your URL"}`)
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ContentTypeJson(w)
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"code": "403", "message": "Method Not Allowed, please check your method"`)
	})
}

func ContentTypeJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func main() {
	initServer()
}
