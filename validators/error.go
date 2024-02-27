package validators

import "reflect"

type validationErr struct {
	Field   string `json:"field,omitempty"`
	Value   any    `json:"value"`
	Message string `json:"message"`
}

func (v *validationErr) Error() string {
	return v.Message
}

func invalidDataType(exp string, got any, lang string) string {
	var t = reflect.TypeOf(got).String()
	if lang == "ar" {
		return "يجب أن تكون البيانات من النوع (" + exp + "), وليس (" + t + ")."
	}
	return "Data should be of type (" + exp + "), not (" + t + ")."
}
