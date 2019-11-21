package cds

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/likexian/gokit/assert"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"reflect"
	"strings"
	"sync/atomic"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"terraform-provider-cds/cds-sdk-go/common/errors"
)

const FILED_SP = "#"

var contextNil context.Context = nil

var logFirstTime = ""
var logAtomicId int64 = 0

// readRetryTimeout is read retry timeout
const readRetryTimeout = 3 * time.Minute

// writeRetryTimeout is write retry timeout
const writeRetryTimeout = 5 * time.Minute

// retryableErrorCode is retryable error code
var retryableErrorCode = []string{
	// client
	"ClientError.NetworkError",
	"ClientError.HttpStatusCodeError",
	// common
	"FailedOperation",
	"TradeUnknownError",
	"RequestLimitExceeded",
	"ResourceInUse",
	"ResourceInsufficient",
	"ResourceUnavailable",
	// cbs
	"ResourceBusy",
}

func init() {
	logFirstTime = fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
}

// getLogId get logId for trace, return a new logId if ctx is nil
func getLogId(ctx context.Context) string {
	if ctx != nil {
		logId, ok := ctx.Value("logId").(string)
		if ok {
			return logId
		}
	}

	return fmt.Sprintf("%s-%d", logFirstTime, atomic.AddInt64(&logAtomicId, 1))
}

// logElapsed log func elapsed time, using in defer
func logElapsed(mark ...string) func() {
	start_at := time.Now()
	return func() {
		log.Printf("[DEBUG] [ELAPSED] %s elapsed %d ms\n", strings.Join(mark, " "), int64(time.Since(start_at)/time.Millisecond))
	}
}

// retryError returns retry error
func retryError(err error, additionRetryableError ...string) *resource.RetryError {
	if isExpectError(err, retryableErrorCode) {
		log.Printf("[CRITAL] Retryable defined error: %v", err)
		return resource.RetryableError(err)
	}

	if len(additionRetryableError) > 0 {
		if isExpectError(err, additionRetryableError) {
			log.Printf("[CRITAL] Retryable addition error: %v", err)
			return resource.RetryableError(err)
		}
	}

	log.Printf("[CRITAL] NonRetryable error: %v", err)

	return resource.NonRetryableError(err)
}

// isExpectError returns whether error is expect error
func isExpectError(err error, expectError []string) bool {
	e, ok := err.(*errors.CdsSDKError)
	if !ok {
		return false
	}

	longCode := e.Code
	if assert.IsContains(expectError, longCode) {
		return true
	}

	if strings.Contains(longCode, ".") {
		shortCode := strings.Split(longCode, ".")[0]
		if assert.IsContains(expectError, shortCode) {
			return true
		}
	}

	return false
}

// writeToFile write data to file
func writeToFile(filePath string, data interface{}) error {
	if strings.HasPrefix(filePath, "~") {
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("Get current user fail,reason %s", err.Error())
		}
		if usr.HomeDir != "" {
			filePath = strings.Replace(filePath, "~", usr.HomeDir, 1)
		}
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("stat old file error,reason %s", err.Error())
	}

	if !os.IsNotExist(err) {
		if fileInfo.IsDir() {
			return fmt.Errorf("old filepath is a dir,can not delete")
		}
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("delete old file error,reason %s", err.Error())
		}
	}

	jsonStr, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("json decode error,reason %s", err.Error())
	}

	return ioutil.WriteFile(filePath, jsonStr, 0422)
}

func CheckNil(object interface{}, fields map[string]string) (nilFields []string) {
	// if object is a pointer, get value which object points to
	object = reflect.Indirect(reflect.ValueOf(object)).Interface()

	for i := 0; i < reflect.TypeOf(object).NumField(); i++ {
		fieldName := reflect.TypeOf(object).Field(i).Name

		if realName, ok := fields[fieldName]; ok {
			if realName == "" {
				realName = fieldName
			}

			if reflect.ValueOf(object).Field(i).IsNil() {
				nilFields = append(nilFields, realName)
			}
		}
	}

	return
}

func BuildTagResourceName(serviceType, resourceType, region, id string) string {
	switch serviceType {
	case "cos":
		return fmt.Sprintf("qcs::%s:%s:uid/:%s/%s", serviceType, region, resourceType, id)

	default:
		return fmt.Sprintf("qcs::%s:%s:uin/:%s/%s", serviceType, region, resourceType, id)
	}
}
