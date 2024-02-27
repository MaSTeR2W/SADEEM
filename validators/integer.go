package validators

import (
	"math"
	"strconv"
)

type Integer struct {
	Field string
	Min   int64
	Max   int64
}

func (i *Integer) GetField() string {
	return i.Field
}

func (i *Integer) Validate(v any, lang string) error {
	// fV: float value
	// json.marshal convert any number to float64
	var fV, ok = v.(float64)

	if !ok || fV != math.Trunc(fV) {
		return &validationErr{
			Field:   i.Field,
			Value:   v,
			Message: invalidDataType("integer", v, lang),
		}
	}

	// iV: integer value
	var iV = int64(fV)

	if iV < i.Min {
		return &validationErr{
			Field:   i.Field,
			Value:   iV,
			Message: smallIntErr(i.Min, lang),
		}
	}

	if iV > i.Max {
		return &validationErr{
			Field:   i.Field,
			Value:   iV,
			Message: bigIntErr(i.Max, lang),
		}
	}

	return nil
}

func smallIntErr(exp int64, lang string) string {
	var sExp = strconv.FormatInt(exp, 10)

	if lang == "ar" {
		return "يجب أن تكون القيمة أكبر من أو تساوي " + sExp + "."
	}

	return "Value should greater than or equal " + sExp + "."
}

func bigIntErr(exp int64, lang string) string {
	var sExp = strconv.FormatInt(exp, 10)

	if lang == "ar" {
		return "يجب أن تكون القيمة أصغر من أو تساوي " + sExp + "."
	}

	return "Value should be less than or equal " + sExp + "."
}
