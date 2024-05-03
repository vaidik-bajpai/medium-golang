package main

import "net/http"

func (app *application) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create blog endpoint"))
}

func (app *application) showBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("show blog endpoint"))
}

func (app *application) listBlogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list blogs endpoint"))
}

func (app *application) updateBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update blog endpoint"))
}

func (app *application) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete blog endpoint"))
}
