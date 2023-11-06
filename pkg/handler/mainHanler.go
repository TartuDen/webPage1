package handler

import (
	"net/http"

	"github.com/TartuDen/HelloWorldWebApp/pkg/render"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "index.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "aboutPage.html")
}
