package errors

// New returns an error that formats as the given text.
func NewKeyInvaildError() error {
	return &KeyInvaildError{s: "key is invaild"}
}

// errorString is a trivial implementation of error.
type KeyInvaildError struct {
	s string
}

func (e *KeyInvaildError) Error() string {
	return e.s
}
