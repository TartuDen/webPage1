package render

import (
	"html/template"
	"net/http"
)

func RenderTmpl(w http.ResponseWriter, tmplName string) {
	tmpl, err := template.ParseFiles("./template/"+tmplName, "./template/body.html")
	if err != nil {
		http.Error(w, "error parsing file"+err.Error(), http.StatusInternalServerError)
	}
	errTmpl := tmpl.Execute(w, nil)
	if errTmpl != nil {
		http.Error(w, "error executing"+errTmpl.Error(), http.StatusInternalServerError)
	}

}
