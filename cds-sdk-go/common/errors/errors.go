package errors

import (
	"fmt"
)

type CdsSDKError struct {
	Code    string
	Message string
	TaskId  string
}

func (e *CdsSDKError) Error() string {
	return fmt.Sprintf("[CdsSDKError] Code=%s, Message=%s, TaskId=%s", e.Code, e.Message, e.TaskId)
}

func NewCdsSDKError(code, message, taskId string) error {
	return &CdsSDKError{
		Code:    code,
		Message: message,
		TaskId:  taskId,
	}
}

func (e *CdsSDKError) GetCode() string {
	return e.Code
}

func (e *CdsSDKError) GetMessage() string {
	return e.Message
}

func (e *CdsSDKError) GetRequestId() string {
	return e.TaskId
}
