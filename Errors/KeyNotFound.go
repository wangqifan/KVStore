package errors

// New returns an error that formats as the given text.
func NewKeyNotFoundError() error {
    return &KeyInvaildError{s : "key is not found"}
}

// errorString is a trivial implementation of error.
type KeyNotFoundError struct {
    s string
}

func (e *KeyNotFoundError) Error() string {
    return e.s
}