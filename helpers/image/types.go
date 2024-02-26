package image

type UnsupportedType struct {
	Field     string `json:"field"`
	MessageAr string `json:"messageAr"`
	MessageEn string `jsin:"messageEn"`
}

func (u *UnsupportedType) String() string {
	return "{\n\t\"field\": \"" + u.Field + "\",\n\t\"messageAr\": \"" + u.MessageAr + "\"\n\t\"messageEn\": \"" + u.MessageEn + "\"\n}"
}

func (u *UnsupportedType) Error() string {
	return u.String()
}
