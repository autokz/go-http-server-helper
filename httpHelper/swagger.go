package httpHelper

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// Swagger adds Swagger documentation handling to the router.
func (r *Router) Swagger(docsPath string) {
	r.swagger("", docsPath)
}

// Swagger adds Swagger documentation handling to the group route.
func (gr *GroupRoute) Swagger(docsPath string) {
	gr.router.swagger(gr.pattern, docsPath)
}

func (r *Router) swagger(pattern, docsPath string) {
	swaggerDocs := http.FileServer(http.Dir(docsPath))
	swaggerDocsHandler := http.StripPrefix(pattern+"/swagger/docs", swaggerDocs).ServeHTTP
	swaggerHandler := httpSwagger.Handler(httpSwagger.URL(pattern + "/swagger/docs/swagger.json"))
	r.HandleFunc("", pattern+"/swagger/docs/", swaggerDocsHandler)
	r.HandleFunc("", pattern+"/swagger/", swaggerHandler)
}
