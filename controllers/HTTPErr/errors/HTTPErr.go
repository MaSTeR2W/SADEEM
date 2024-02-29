package errors

type HTTP400Err struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (h *HTTP400Err) Error() string {
	return h.Message
}

type HTTP404Err struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func (h *HTTP404Err) Error() string {
	return h.Message
}

type HTTP401Err struct {
	Message string `json:"message"`
}

func (h *HTTP401Err) Error() string {
	return h.Message
}

type HTTP403Err struct {
	Message string `json:"message"`
}

func (h *HTTP403Err) Error() string {
	return h.Message
}
