package security

import (
	"github.com/MaSTeR2W/SADEEM/validators"
)

var password = &validators.String{
	Field:     "password",
	NotNil:    true,
	MinLength: 8,
	MaxLength: 16,
}

var email = &validators.Email{
	Field:  "email",
	NotNil: true,
}

var LoginValidator = validators.Object{
	Field:    "body",
	Required: []string{"email", "password"},
	NotNil:   true,

	Props: []validators.Validator{password, email},
}
