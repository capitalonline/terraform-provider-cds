package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"terraform-provider-cds/cds-sdk-go/common/errors"

	//"log"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code   string `json:"Code"`
			TaskId string `json:"TaskId"`
		} `json:"Error" omitempty`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

type DeprecatedAPIErrorResponse struct {
	Code    string `json:"Code"`
	Message string `json:"message"`
	TaskId  string `json:"TaskId"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewCdsSDKError("ClientError.ParseJsonError", msg, "")
	}
	//if resp.Response.Error.Code != "" {
	//	return errors.NewCdsSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	//}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewCdsSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != "Success" {
		return errors.NewCdsSDKError(deprecated.Code, deprecated.Message, "")
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, response Response, r Request) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.NewCdsSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		fmt.Println(r.GetAction())
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.NewCdsSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	//log.Printf("[DEBUG] Response Body=%s", body)
	err = response.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewCdsSDKError("ClientError.ParseJsonError", msg, "")
	}
	return
}
