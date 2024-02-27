package validators

type validator interface {
	Validate(any, string) error
	GetField() string
}
