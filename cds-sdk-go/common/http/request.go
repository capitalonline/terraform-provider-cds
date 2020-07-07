package common

import (
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	POST = "POST"
	GET  = "GET"

	RootDomain = "cdsapi.capitalonline.net"
	Path       = "/"
)

type Request interface {
	GetAction() string
	GetBodyReader() io.Reader
	GetDomain() string
	GetHttpMethod() string
	GetParams() map[string]string
	GetPath() string
	GetService() string
	GetUrl() string
	GetVersion() string
	SetDomain(string)
	SetHttpMethod(string)
	SetPath(string)
}

type BaseRequest struct {
	httpMethod string
	domain     string
	path       string
	params     map[string]string
	formParams map[string]string

	service string
	version string
	action  string
}

func (r *BaseRequest) GetAction() string {
	return r.action
}

func (r *BaseRequest) GetHttpMethod() string {
	return r.httpMethod
}

func (r *BaseRequest) GetParams() map[string]string {
	return r.params
}

func (r *BaseRequest) GetPath() string {
	return r.path
}

func (r *BaseRequest) SetPath(path string) {
	r.path = strings.ToLower(path)
}

func (r *BaseRequest) GetDomain() string {
	return r.domain
}

func (r *BaseRequest) SetDomain(domain string) {
	r.domain = domain
}

func (r *BaseRequest) SetHttpMethod(method string) {
	switch strings.ToUpper(method) {
	case POST:
		{
			r.httpMethod = POST
		}
	case GET:
		{
			r.httpMethod = GET
		}
	default:
		{
			r.httpMethod = GET
		}
	}
}

func (r *BaseRequest) GetService() string {
	return r.service
}

func (r *BaseRequest) GetUrl() string {
	if r.httpMethod == GET {
		return "http://" + r.domain + r.path
	} else if r.httpMethod == POST {
		return "http://" + r.domain + r.path
	} else {
		return ""
	}
}

func (r *BaseRequest) GetVersion() string {
	return r.version
}

func GetUrlQueriesEncoded(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		if value != "" && key != "SecretId" {
			values.Add(key, value)
		}
	}
	return values.Encode()
}

func (r *BaseRequest) GetBodyReader() io.Reader {
	if r.httpMethod == POST {
		s := GetUrlQueriesEncoded(r.params)
		return strings.NewReader(s)
	} else {
		return strings.NewReader("")
	}
}

func (r *BaseRequest) Init() *BaseRequest {
	r.domain = ""
	r.path = Path
	r.params = make(map[string]string)
	r.formParams = make(map[string]string)
	return r
}

func (r *BaseRequest) WithApiInfo(service, version, action string) *BaseRequest {
	r.service = service
	r.version = version
	r.action = action
	return r
}

func GetServiceDomain(service string) (domain string) {
	domain = RootDomain + "/" + strings.ToLower(service)
	return
}

func ConstructParams(req Request) (err error) {
	value := reflect.ValueOf(req).Elem()
	err = flatStructure(value, req, "")
	//log.Printf("[DEBUG] params=%s", req.GetParams())
	return
}

func CompleteCdsParams(ak string, req Request) {
	params := req.GetParams()
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	uuidStr := uuid.New()
	params["Action"] = req.GetAction()
	params["AccessKeyId"] = ak
	params["SignatureNonce"] = uuidStr.String()
	params["SignatureVersion"] = "1.0"
	params["Timestamp"] = timestamp
	if req.GetVersion() != "" {
		params["Version"] = req.GetVersion()
	} else {
		params["Version"] = "2019-08-08"
	}
	return
}

//func CdsConstructParams(ak,sk string, request Request) (signatureUrl string, err error) {
//	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
//	uuidStr, _ := uuid.NewV4()
//	var paramsKeys sort.StringSlice
//
//	data := map[string]string{
//		"Action":           request.GetAction(),
//		"AccessKeyId":      ak,
//		"SignatureMethod":  "HMAC-SHA1",
//		"SignatureNonce":   uuidStr.String(),
//		"SignatureVersion": "1.0",
//		"Timestamp":        timestamp,
//		"Version":          "2019-08-08",
//	}
//
//	if request.GetHttpMethod()==GET && request.GetParams() != nil {
//		for k, v := range request.GetParams() {
//			data[k] = v
//		}
//	}
//
//	for k, _ := range data {
//		paramsKeys = append(paramsKeys, k)
//	}
//	sort.Sort(paramsKeys)
//
//	var urlStr string
//	for _, k := range paramsKeys {
//		urlStr += "&" + percentEncode(k) + "=" + percentEncode(data[k])
//	}
//	urlStr = request.GetHttpMethod() + "&%2F&" + percentEncode(urlStr[1:])
//	h := hmac.New(sha1.New, []byte(sk))
//	h.Write([]byte(urlStr))
//	encodeString := base64.StdEncoding.EncodeToString(h.Sum(nil))
//
//	data["Signature"] = encodeString
//
//	val := url.Values{}
//	for k, v := range data {
//		val.Add(k, v)
//	}
//	signatureUrl = val.Encode()
//	return signatureUrl, nil
//}
//
//func percentEncode(str string) string {
//	str = url.QueryEscape(str)
//	strings.Replace(str, "+", "%20", -1)
//	strings.Replace(str, "*", "%2A", -1)
//	strings.Replace(str, "%7E", "~", -1)
//	return str
//}

func flatStructure(value reflect.Value, request Request, prefix string) (err error) {
	//log.Printf("[DEBUG] reflect value: %v", value.Type())
	valueType := value.Type()
	for i := 0; i < valueType.NumField(); i++ {
		tag := valueType.Field(i).Tag
		nameTag, hasNameTag := tag.Lookup("name")
		if !hasNameTag {
			continue
		}
		field := value.Field(i)
		kind := field.Kind()
		if kind == reflect.Ptr && field.IsNil() {
			continue
		}
		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}
		key := prefix + nameTag
		if kind == reflect.String {
			s := field.String()
			if s != "" {
				request.GetParams()[key] = s
			}
		} else if kind == reflect.Bool {
			request.GetParams()[key] = strconv.FormatBool(field.Bool())
		} else if kind == reflect.Int || kind == reflect.Int64 {
			request.GetParams()[key] = strconv.FormatInt(field.Int(), 10)
		} else if kind == reflect.Uint || kind == reflect.Uint64 {
			request.GetParams()[key] = strconv.FormatUint(field.Uint(), 10)
		} else if kind == reflect.Float64 {
			request.GetParams()[key] = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		} else if kind == reflect.Slice {
			list := value.Field(i)
			for j := 0; j < list.Len(); j++ {
				vj := list.Index(j)
				key := prefix + nameTag + "." + strconv.Itoa(j)
				kind = vj.Kind()
				if kind == reflect.Ptr && vj.IsNil() {
					continue
				}
				if kind == reflect.Ptr {
					vj = vj.Elem()
					kind = vj.Kind()
				}
				if kind == reflect.String {
					request.GetParams()[key] = vj.String()
				} else if kind == reflect.Bool {
					request.GetParams()[key] = strconv.FormatBool(vj.Bool())
				} else if kind == reflect.Int || kind == reflect.Int64 {
					request.GetParams()[key] = strconv.FormatInt(vj.Int(), 10)
				} else if kind == reflect.Uint || kind == reflect.Uint64 {
					request.GetParams()[key] = strconv.FormatUint(vj.Uint(), 10)
				} else if kind == reflect.Float64 {
					request.GetParams()[key] = strconv.FormatFloat(vj.Float(), 'f', -1, 64)
				} else {
					if err = flatStructure(vj, request, key+"."); err != nil {
						return
					}
				}
			}
		} else {
			if err = flatStructure(reflect.ValueOf(field.Interface()), request, prefix+nameTag+"."); err != nil {
				return
			}
		}
	}
	return
}
