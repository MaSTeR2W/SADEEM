package user

import "github.com/MaSTeR2W/SADEEM/validators"

var name = &validators.String{
	Field:     "name",
	NotNil:    true,
	MinLength: 3,
	MaxLength: 50,
}

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

var userType = &validators.String{
	Field:  "userType",
	NotNil: true,
	Enum:   []string{"normal", "manager"},
}

var CreationValidator = validators.Object{
	Field: "body",
	Required: []string{
		"name",
		"email",
		"password",
		"userType",
		"image",
	},
	NotNil: true,
	Props: []validators.Validator{
		name,
		password,
		email,
		userType,
	},
}

var UpdateValidator = validators.Object{
	Field:  "body",
	NotNil: true,
	OneOf: []string{
		"name",
		"email",
		"password",
		"image",
	},
	Props: []validators.Validator{
		name,
		password,
		email,
		userType,
	},
}
