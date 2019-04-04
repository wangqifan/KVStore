package errors

// New returns an error that formats as the given text.
func NewValueInvaildError() error {
    return &ValueInvaildError{s : "value is invaild"}
}

// errorString is a trivial implementation of error.
type ValueInvaildError struct {
    s string
}

func (e *ValueInvaildError) Error() string {
    return e.s
}