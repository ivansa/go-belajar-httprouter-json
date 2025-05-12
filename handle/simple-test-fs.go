package handle

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"net/http"
)

//go:embed resources/test-fs
var resources embed.FS

// http://localhost:8080/test-router-fs/sample.json
func RegisterTestSimpleRoutesWithFS(router *httprouter.Router) {
	dir, _ := fs.Sub(resources, "resources/test-fs")
	router.ServeFiles("/test-router-fs/*filepath", http.FS(dir))
}
