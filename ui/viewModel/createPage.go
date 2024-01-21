package viewmodel

import "github.com/dchupp/snippetbox/internal/validator"

type CreateSnippetForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}
