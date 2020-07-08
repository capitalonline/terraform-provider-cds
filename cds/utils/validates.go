package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/athom/goset"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func validateNameRegex(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if _, err := regexp.Compile(value); err != nil {
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid regular expression: %s",
			k, err))
	}
	return
}

func validateNotEmpty(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) == 0 {
		errors = append(errors, fmt.Errorf("should not use empty string"))
	}
	return
}

func validateInstanceType(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	words := strings.Split(value, ".")
	if len(words) <= 1 {
		errors = append(errors, fmt.Errorf("invalid instance_type: %v, should be like S1.SMALL1", value))
		return
	}
	return
}

// validateCIDRNetworkAddress ensures that the string value is a valid CIDR that
// represents a network address - it adds an error otherwise
func validateCIDRNetworkAddress(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	_, ipnet, err := net.ParseCIDR(value)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q must contain a valid CIDR, got error parsing: %s", k, err))
		return
	}
	if ipnet == nil || value != ipnet.String() {
		errors = append(errors, fmt.Errorf("%q must contain a valid network CIDR, expected %q, got %q", k, ipnet, value))
	}
	return
}

func validateIp(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	ip := net.ParseIP(value)
	if ip == nil {
		errors = append(errors, fmt.Errorf("%q must contain a valid IP", k))
	}
	return
}

// NOTE not exactly strict, but ok for now
func validateIntegerInRange(min, max int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < min {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d: %d", k, min, value))
		}
		if value > max {
			errors = append(errors, fmt.Errorf(
				"%q cannot be higher than %d: %d", k, max, value))
		}
		return
	}
}

func validateIntegerMin(min int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < min {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d: %d", k, min, value))
		}
		return
	}
}

func ValidateStringLengthInRange(min, max int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := len(v.(string))
		if value < min {
			errors = append(errors, fmt.Errorf(
				"length of %q cannot be lower than %d: %d", k, min, value))
		}
		if value > max {
			errors = append(errors, fmt.Errorf(
				"length of %q cannot be higher than %d: %d", k, max, value))
		}
		return
	}
}

func validateKeyPairName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 25 || len(value) == 0 {
		errors = append(errors, fmt.Errorf("invalid key pair: %v, size too long or too short", value))
	}

	pattern := `^[a-zA-Z0-9_]+$`
	if match, _ := regexp.Match(pattern, []byte(value)); !match {
		errors = append(errors, fmt.Errorf("invalid key pair: %v, wrong format", value))
	}
	return
}

func validateAllowedStringValueIgnoreCase(ss []string) schema.SchemaValidateFunc {

	var upperStrs = make([]string, len(ss))
	for index, value := range ss {
		upperStrs[index] = strings.ToUpper(value)
	}
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !goset.IsIncluded(upperStrs, strings.ToUpper(value)) {
			errors = append(errors, fmt.Errorf("%q must contain a valid string value should in array %#v, got %q", k, ss, value))
		}
		return
	}
}

func validateAllowedStringValue(ss []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !goset.IsIncluded(ss, value) {
			errors = append(errors, fmt.Errorf("%q must contain a valid string value should in array %#v, got %q", k, ss, value))
		}
		return
	}
}

func validatePort(v interface{}, k string) (ws []string, errors []error) {
	value := 0
	switch t := v.(type) {
	case string:
		value, _ = strconv.Atoi(t)
	case int:
		value = t
	default:
		errors = append(errors, fmt.Errorf("%q data type error ", k))
		return
	}
	if value < 1 || value > 65535 {
		errors = append(errors, fmt.Errorf("%q must be a valid port between 1 and 65535", k))
	}
	return
}

func validateMysqlPassword(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 64 || len(value) < 8 {
		errors = append(errors, fmt.Errorf("invalid password, len(password) must between 8 and 64,%s", value))
	}
	var match = make(map[string]bool)
	if strings.ContainsAny(value, "_+-&=!@#$%^*()") {
		match["alien"] = true
	}
	for i := 0; i < len(value); i++ {
		if len(match) >= 2 {
			break
		}
		if value[i] >= '0' && value[i] <= '9' {
			match["number"] = true
			continue
		}
		if (value[i] >= 'a' && value[i] <= 'z') || (value[i] >= 'A' && value[i] <= 'Z') {
			match["letter"] = true
			continue
		}
	}
	if len(match) < 2 {
		errors = append(errors, fmt.Errorf("invalid password, contains at least letters, Numbers, and characters(_+-&=!@#$%%^*()),%s", value))
	}
	return
}

func validateAllowedIntValue(ints []int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if !goset.IsIncluded(ints, value) {
			errors = append(errors, fmt.Errorf("%q must contain a valid string value should in array %#v, got %q", k, ints, value))
		}
		return
	}
}

// Only support lowercase letters, numbers and "-". It cannot be longer than 40 characters.
func validateCosBucketName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 40 || len(value) < 0 {
		errors = append(errors, fmt.Errorf("invalid bucket name: %v, size too long or too short", value))
	}

	pattern := `^[a-z0-9-]+$`
	if match, _ := regexp.Match(pattern, []byte(value)); !match {
		errors = append(errors, fmt.Errorf("invalid bucket name: %v, wrong format: only support lowercase letters, numbers and -", value))
	}
	return
}

func validateCosBucketLifecycleTimestamp(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	_, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", value))
	if err != nil {
		errors = append(errors, fmt.Errorf(
			"%q cannot be parsed as RFC3339 Timestamp Format", value))
	}

	return
}

func validateAsConfigPassword(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 16 || len(value) < 8 {
		errors = append(errors, fmt.Errorf("invalid password, len(password) must between 8 and 16,%s", value))
	}
	var match = make(map[string]bool)
	if strings.ContainsAny(value, "()~!@#$%^&*-+={}[]:;',.?/") {
		match["alien"] = true
	}
	for i := 0; i < len(value); i++ {
		if len(match) >= 2 {
			break
		}
		if value[i] >= '0' && value[i] <= '9' {
			match["number"] = true
			continue
		}
		if (value[i] >= 'a' && value[i] <= 'z') || (value[i] >= 'A' && value[i] <= 'Z') {
			match["letter"] = true
			continue
		}
	}
	if len(match) < 2 {
		errors = append(errors, fmt.Errorf("invalid password, contains at least letters, Numbers, and characters(_+-&=!@#$%%^*()),%s", value))
	}
	return
}

func validateAsScheduleTimestamp(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	_, err := time.Parse(time.RFC3339, value)
	if err != nil {
		errors = append(errors, fmt.Errorf(
			"%q cannot be parsed as RFC3339 Timestamp Format", value))
	}
	return
}

func validateStringPrefix(prefix string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !strings.HasPrefix(value, prefix) {
			errors = append(errors, fmt.Errorf("%s doesn't have preifx %s", k, prefix))
		}
		return
	}
}

func validateCidrIp(v interface{}, k string) (ws []string, errs []error) {
	if _, err := validateIp(v, k); len(err) == 0 {
		return
	}

	if _, err := validateCIDRNetworkAddress(v, k); len(err) != 0 {
		errs = append(errs, fmt.Errorf("%s %v is not valid IP address or valid CIDR IP address",
			k, v))
	}
	return
}
