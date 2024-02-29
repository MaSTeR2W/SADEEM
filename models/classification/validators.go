package classification

import "github.com/MaSTeR2W/SADEEM/validators"

var name = &validators.String{
	Field:     "name",
	NotNil:    true,
	MinLength: 1,
	MaxLength: 20,
}

var enabled = &validators.Bool{
	Field:  "enabled",
	NotNil: true,
}

var CreationValidator = validators.Object{
	Field:    "body",
	NotNil:   true,
	Required: []string{"name", "enabled"},
	Props: []validators.Validator{
		name,
		enabled,
	},
}

var UpdateValidator = validators.Object{
	Field:  "body",
	NotNil: true,
	OneOf:  []string{"name", "enabled"},
	Props: []validators.Validator{
		name,
		enabled,
	},
}
