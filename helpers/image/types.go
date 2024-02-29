package image

type UnsupportedType struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (u *UnsupportedType) String() string {
	return "{\n\t\"field\": \"" + u.Field + "\"\n\t\"messageEn\": \"" + u.Message + "\"\n}"
}

func (u *UnsupportedType) Error() string {
	return u.String()
}
