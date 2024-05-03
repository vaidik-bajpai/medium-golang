package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", app.healthCheckHandler)

	mux.HandleFunc("POST /v1/blogs", app.createBlogHandler)
	mux.HandleFunc("GET /v1/blogs/{id}", app.showBlogHandler)
	mux.HandleFunc("GET /v1/blogs", app.listBlogsHandler)
	mux.HandleFunc("PUT /v1/blogs/{id}", app.updateBlogHandler)
	mux.HandleFunc("DELETE /v1/blogs/{id}", app.deleteBlogHandler)

	return mux
}
