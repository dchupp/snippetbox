package ui

import (
	"time"

	"github.com/dchupp/snippetbox/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type TemplateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Flash           string
	IsAuthenticated bool   // Add an IsAuthenticated field to the templateData struct.
	CSRFToken       string // Add a CSRFToken field.

}
type FormFieldErrors struct {
	FieldName   string
	FieldErrors []string
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func HumanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
// var functions = template.FuncMap{
// 	"humanDate": HumanDate,
// }

// func NewTemplateCache() (map[string]*template.Template, error) {
// 	cache := map[string]*template.Template{}

// 	// Use fs.Glob() to get a slice of all filepaths in the ui.Files embedded
// 	// filesystem which match the pattern 'html/pages/*.tmpl'. This essentially
// 	// gives us a slice of all the 'page' templates for the application, just
// 	// like before.
// 	pages, err := fs.Glob(Files, "html/pages/*.go.tmpl")
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)

// 		// Create a slice containing the filepath patterns for the templates we
// 		// want to parse.
// 		patterns := []string{
// 			"html/base.go.tmpl",
// 			"html/partials/*.go.tmpl",
// 			page,
// 		}

// 		// Use ParseFS() instead of ParseFiles() to parse the template files
// 		// from the ui.Files embedded filesystem.
// 		ts, err := template.New(name).Funcs(functions).ParseFS(Files, patterns...)
// 		if err != nil {
// 			return nil, err
// 		}

// 		cache[name] = ts
// 	}

// 	return cache, nil
// }