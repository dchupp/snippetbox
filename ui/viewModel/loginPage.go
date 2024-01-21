package viewmodel

import "github.com/dchupp/snippetbox/internal/validator"

type UserLoginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}
