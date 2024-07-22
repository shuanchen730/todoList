package ecode

import (
	"fmt"

	"github.com/pkg/errors"
)

// 參考：https://segmentfault.com/q/1010000019676525
var (
	codes map[string]Error
)

type detail struct {
	RedirectCode    string
	RedirectMessage string
	RedirectDetails interface{}
}

type Error struct {
	code       string
	message    string
	innerError string
	detail     detail
}

// add only inner error
func add(code string, msg string) Error {
	if codes == nil {
		codes = make(map[string]Error)
	}
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("ecode: %s already exist", code))
	}
	codes[code] = Error{
		code:       code,
		message:    msg,
		innerError: "default error.",
		detail: detail{
			RedirectCode:    "",
			RedirectMessage: "",
		},
	}
	return codes[code]
}

type Errors interface {
	// sometimes Error return Code in string form
	Error() string
	// Code get error code.
	Code() string
	// Message get code message.
	Message() string
	// Equal for compatible.
	Equal(error) bool
	// Reload Message
	Reload(string) Error

	SetInnerError(error []string) Error

	SetRedirectDetails(details interface{}) Error

	SetDetail(code string, msg string, details interface{}) Error

	GetDetail() (code string, message string, details interface{})
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Equal(err error) bool {
	return Equal(err, e)
}

func Equal(err error, e Error) bool {
	return Cause(err).Code() == e.Code()
}

func (e Error) Reload(message string) Error {
	e.message = message
	return e
}

func (e Error) SetInnerError(error []string) Error {
	if len(error) > 0 {
		for _, errorString := range error {
			e.innerError = e.innerError + " " + errorString
		}
	}
	return e
}

func (e Error) SetRedirectDetails(details interface{}) Error {
	e.detail.RedirectDetails = details
	return e
}

func (e Error) SetDetail(code string, message string, details interface{}) Error {
	e.detail.RedirectCode = code
	e.detail.RedirectMessage = message
	e.detail.RedirectDetails = details
	return e
}

func (e Error) GetDetail() (code string, message string, details interface{}) {
	return e.detail.RedirectCode, e.detail.RedirectMessage, e.detail.RedirectDetails
}

func String(e string) Error {
	if e == "" {
		return Error{}
	}
	return Error{
		code: "210299999", message: e,
	}
}

func Cause(err error, innerError ...string) Errors {
	if err == nil {
		return nil
	}
	if ec, ok := errors.Cause(err).(Errors); ok {
		ec = ec.SetInnerError(innerError)
		return ec
	}
	return String(err.Error())
}

func CauseWithDetail(err error, redirectCode, redirectMsg string) Errors {
	if err == nil {
		return nil
	}
	if ec, ok := errors.Cause(err).(Errors); ok {
		ec = ec.SetDetail(redirectCode, redirectMsg, nil)
		return ec
	}
	return String(err.Error())
}
