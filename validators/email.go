package validators

import (
	"strings"
)

// RFC 5322
type Email struct {
	Field string
}

func (e *Email) GetField() string {
	return e.Field
}

func (e *Email) Validate(v any, lang string) error {
	var email, ok = v.(string)

	if !ok {
		return &validationErr{
			Field:   e.Field,
			Value:   v,
			Message: invalidDataType("string", v, lang),
		}
	}

	var emailLen = len(email)

	if emailLen < 5 {
		return &validationErr{
			Field:   e.Field,
			Value:   email,
			Message: "",
		}
	}

	if emailLen > 320 {
		return &validationErr{
			Field:   e.Field,
			Value:   email,
			Message: "",
		}
	}

	var parts = strings.Split(email, "@")

	var partsLen = len(parts)

	if partsLen == 1 {
		return &validationErr{
			Field:   e.Field,
			Value:   email,
			Message: "",
		}
	}

	if partsLen > 2 {
		return &validationErr{
			Field:   e.Field,
			Value:   email,
			Message: "",
		}
	}

	var err = IsLocalPartValid(e.Field, email, parts[0])

	if err != nil {
		return err
	}

	return IsDomainNameValid(e.Field, email, parts[1])

}

var domainAllowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-"

func IsDomainNameValid(field, email, domain string) error {
	var l = len(domain)

	if l < 1 {
		return &validationErr{
			Field:   field,
			Value:   email,
			Message: "",
		}
	}

	if l > 255 {
		return &validationErr{
			Field:   field,
			Value:   email,
			Message: "",
		}
	}

	for _, r := range domain {
		if !strings.ContainsRune(domainAllowedChars, r) {
			return &validationErr{
				Field:   field,
				Value:   email,
				Message: "",
			}
		}
	}

	// check mx record for domain
	/* mxRcds, err := net.LookupMX(domain)

	if err != nil || len(mxRcds) == 0 {
		return &validationErr{
			Field:   field,
			Value:   email,
			Message: "",
		}
	} */

	return nil
}

var allowedLocalPartChars = domainAllowedChars + "_.-+!#$%&'*/=?`^{}[]|~"

func IsLocalPartValid(field, email, localPart string) error {
	var l = len(localPart)

	if l == 0 {
		return &validationErr{
			Field:   field,
			Value:   email,
			Message: "",
		}
	}

	if l > 64 {
		return &validationErr{
			Field:   field,
			Value:   email,
			Message: "",
		}
	}

	for _, r := range localPart {
		if !strings.ContainsRune(allowedLocalPartChars, r) {
			return &validationErr{
				Field:   field,
				Value:   email,
				Message: "",
			}
		}
	}

	return nil
}
