package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
	_http "terraform-provider-cds/cds-sdk-go/common/http"
)

const (
	SHA256 = "HmacSHA256"
	SHA1   = "HMAC-SHA1"
)

func Sign(s, secretKey, method string) string {
	hashed := hmac.New(sha1.New, []byte(secretKey))
	if method == SHA256 {
		hashed = hmac.New(sha256.New, []byte(secretKey))
	}
	hashed.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func signRequest(request _http.Request, credential *Credential, method string) (err error) {
	if method != SHA256 {
		method = SHA1
	}
	checkAuthParams(request, credential, method)
	s := getStringToSign(request)
	signature := Sign(s, credential.SecretKey, method)
	request.GetParams()["Signature"] = signature
	return
}

func checkAuthParams(request _http.Request, credential *Credential, method string) {
	params := request.GetParams()
	//credentialParams := credential.GetCredentialParams()
	//for key, value := range credentialParams {
	//	params[key] = value
	//}
	params["SignatureMethod"] = method
	delete(params, "Signature")
}

func getStringToSign(request _http.Request) string {
	method := request.GetHttpMethod()
	params := request.GetParams()
	var paramsKeys sort.StringSlice
	for k, _ := range params {
		paramsKeys = append(paramsKeys, k)
	}
	sort.Sort(paramsKeys)
	var urlStr string
	for _, k := range paramsKeys {
		urlStr += "&" + percentEncode(k) + "=" + percentEncode(params[k])
	}
	urlStr = method + "&%2F&" + percentEncode(urlStr[1:])
	return urlStr
}

func percentEncode(str string) string {
	str = url.QueryEscape(str)
	strings.Replace(str, "+", "%20", -1)
	strings.Replace(str, "*", "%2A", -1)
	strings.Replace(str, "%7E", "~", -1)
	return str
}
