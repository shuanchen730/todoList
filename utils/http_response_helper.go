package utils

import (
	"todoList/ecode"
	"todoList/entities/delivery"
)

func MakeECodeResponse(eCode ecode.Errors, message ...string) delivery.HttpResponse {
	errResp := delivery.HttpResponse{}
	errResp.Result = 0
	errResp.ErrorCode = eCode.Code()
	errResp.Message = eCode.Message()
	if len(message) > 0 {
		for _, msg := range message {
			errResp.Message = errResp.Message + " " + msg
		}
	}
	return errResp
}
