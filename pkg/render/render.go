package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// tc is a map that stores template templates by string keys.
var tc = make(map[string]*template.Template)

// RenderTmpl renders a template to the given HTTP response writer.
func RenderTmpl(w http.ResponseWriter, t string) {
	// Declare variables for the template and error.
	var tmpl *template.Template
	var err error

	// Check if the template is already in the cache.
	_, inMap := tc[t]
	if !inMap {
		// If not in cache, create and cache a new template.
		fmt.Println("We create and cache a new template")
		err = createTemplatCache(t)
		if err != nil {
			http.Error(w, "error Executing"+err.Error(), http.StatusInternalServerError)
		}
	} else {
		// If the template is in the cache, use the cached template.
		fmt.Println("We use the cached template")
	}

	// Retrieve the template from the cache.
	tmpl = tc[t]

	// Execute the template and handle any execution errors.
	errExecute := tmpl.Execute(w, nil)
	if errExecute != nil {
		http.Error(w, "error Executing"+errExecute.Error(), http.StatusInternalServerError)
	}
}

// createTemplatCache creates and caches a template for the specified key.
func createTemplatCache(t string) error {
	// List of template files to be used in the template.
	templates := []string{
		fmt.Sprintf("./template/%s", t),
		"./template/body.html",
	}

	// Parse the template files.
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Cache the parsed template for the specified key.
	tc[t] = tmpl
	return nil
}
