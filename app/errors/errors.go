package app

import (
	"errors"
	"net/http"
	"io"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

const (
	InvalidRequestError = "invalid_request_error"
	ApiError            = "api_error"
)

var (
	MGTRecordNotFound = errors.New("Record not found.")
	MGTUnrecognizedURL = errors.New("Unrecognized request URL.")
	MGTMethodNotAllowed = errors.New("Method not allowed.")
	MGTInvalidParams = errors.New("Parameters is invalid")
	MGTEmptyParams = errors.New("Parameters cannot be empty.")
)

const ()

type MGTErrorJSON struct {
	Error *MGTError `json:"error"`
}

type MGTError struct {
	Type        string 						`json:"type"`
	Description string 						`json:"description"`
	HttpCode    int    						`json:"-"`
	Details     []*revel.ValidationError 	`json:"details,omitempty"`
}

func (e *MGTError) Error() string {
	return e.Description
}

func (e *MGTError) ToJSON() *MGTErrorJSON {
	return &MGTErrorJSON{e}
}

// need more spesific error, we still dont handle HTTP verbs error in route
func NewMGTError(err error, details ...interface{}) *MGTError {
	switch err {
	case gorm.RecordNotFound:
		return &MGTError{InvalidRequestError, MGTRecordNotFound.Error(), http.StatusBadRequest,nil}
	case MGTUnrecognizedURL:
		return &MGTError{InvalidRequestError, err.Error(), http.StatusNotFound, nil}
	case MGTMethodNotAllowed:
		return &MGTError{InvalidRequestError, err.Error(), http.StatusMethodNotAllowed, nil}
	case MGTInvalidParams:
		mgtError := &MGTError{InvalidRequestError, err.Error(), http.StatusBadRequest, nil}
		if details != nil  {
			if valErrors, ok := details[0].([]*revel.ValidationError); ok {
				mgtError.Details = valErrors
			}
		}
		return mgtError
	case io.EOF:
		return &MGTError{InvalidRequestError, MGTEmptyParams.Error(), http.StatusBadRequest, nil}
	default:
		return &MGTError{ApiError, err.Error(), http.StatusInternalServerError, nil}
	}
}
