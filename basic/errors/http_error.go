package errors

import (
	"fmt"
	"io"
	"strings"
)

func NewHTTPError(code ErrorCode, status int, err ...error) *HTTPError {
	e := &HTTPError{
		Code:       (code),
		HTTPStatus: status,
	}
	if len(err) > 0 {
		e.error = err[0]
		e.Message = e.Error()
	}
	return e
}

type HTTPError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	HTTPStatus int                    `json:"-"`
	Log        bool                   `json:"-"`
	Data       map[string]interface{} `json:"-"`
	error      error
}

type ErrorData map[string]interface{}

func (herr *HTTPError) WithData(data ErrorData) *HTTPError {
	if herr.Data == nil {
		herr.Data = map[string]interface{}{}
	}
	for k, v := range data {
		herr.Data[k] = v
		herr.Message = strings.ReplaceAll(herr.Message, "{"+k+"}", fmt.Sprintf("%v", v))
	}
	return herr
}

func (herr *HTTPError) Cause() error {
	return herr.error
}

func (herr *HTTPError) Error() string {
	if herr.error == nil {
		return herr.Message
	}
	return herr.Message + " - " + herr.error.Error()
}

func (w *HTTPError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.Message)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

func (herr *HTTPError) Get(key string) (interface{}, bool) {
	var value interface{}
	switch key {
	case "code":
		value = herr.Code
	case "message":
		value = herr.Message
	case "http_status":
		value = herr.HTTPStatus
	}
	if value != nil {
		return value, true
	}
	return nil, false
}

func (herr *HTTPError) Set(key string, value interface{}) {
	switch key {
	case "code":
		herr.Code = ErrorCode(value.(int))
	case "message":
		herr.Message = value.(string)
	case "http_status":
		herr.HTTPStatus = value.(int)
	}
}
