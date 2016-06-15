package oxford

import (
	"errors"
	"net/http"
)

type OxofordError error
type SpeakError error

type OxfordError struct {
	HttpStatus int
	error
}

func NewError(httpStatus int, err error) OxfordError {
	return OxfordError{httpStatus, err}
}

func (o *OxfordError) setError(err error) {
	o.error = err
}

var (
	errBadRequest    = NewError(http.StatusBadRequest, errors.New("BadRequest"))
	errUnknownLocale = NewError(http.StatusBadRequest, errors.New("UnknownLocale"))
)

func IsSpeakError(err error) bool {
	_, ok := err.(SpeakError)
	return ok
}
