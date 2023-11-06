package render

import (
	"fmt"
	"html/template"
	"net/http"
)

var tc = make(map[string]*template.Template)

func RenderTmpl(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		fmt.Println("We create and cache new template")
		err = createTemplatCache(t)
		if err != nil {
			http.Error(w, "error Executing"+err.Error(), http.StatusInternalServerError)
		}

	} else {
		fmt.Println("We use cached template")
	}
	tmpl = tc[t]
	errExecute := tmpl.Execute(w, nil)
	if errExecute != nil {
		http.Error(w, "error Executing"+errExecute.Error(), http.StatusInternalServerError)

	}
}

func createTemplatCache(t string) error {
	templates := []string{
		fmt.Sprintf("./template/%s", t),
		"./template/body.html",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
