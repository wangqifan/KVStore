package errors

// New returns an error that formats as the given text.
func NewScanParameterInvaildError() error {
	return &ScanParameterInvaildError{s: "the parameter of scan is invaild"}
}

// errorString is a trivial implementation of error.
type ScanParameterInvaildError struct {
	s string
}

func (e *ScanParameterInvaildError) Error() string {
	return e.s
}
