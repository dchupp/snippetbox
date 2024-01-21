package viewmodel

import "github.com/dchupp/snippetbox/internal/validator"

type UserSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}
