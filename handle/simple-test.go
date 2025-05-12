package handle

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"code": "%d", "message": "%s"}`, code, message)
}

func RegisterTestSimpleRoutes(router *httprouter.Router) {
	router.GET("/test-router", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		JsonResponse(w, 200, "Simple Test Router")
	})

	//httprouter.Params -> path variable
	router.GET("/test-router/name/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		JsonResponse(w, 200, "Simple Test Router, Hello "+name)
	})

	router.GET("/test-router/name/:name/with-age/:age", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		age := ps.ByName("age")
		JsonResponse(w, 200, "Simple Test Router, Hello "+name+"! You are "+age+" years old!")
	})

	router.GET("/test-router/path/*path", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		path := ps.ByName("path")
		JsonResponse(w, 200, "Simple Test Router, Path : "+path)
	})
}
