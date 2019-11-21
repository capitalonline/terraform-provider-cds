package utils

import (
	"errors"
	"reflect"
	"strconv"
	"unicode"
)

func Mapstructure(m map[string]interface{}, v interface{}) error {

	var (
		objType  reflect.Type
		objValue reflect.Value

		field      reflect.StructField
		fieldValue reflect.Value

		fieldName string
	)
	if v == nil {
		return errors.New("result is nil")
	}

	objType = reflect.TypeOf(v)
	objValue = reflect.ValueOf(v)

	if objType.Kind() == reflect.Ptr {
		if objValue.IsNil() {
			return errors.New("result must be addressable (a pointer)")
		}
		objType = objType.Elem()
		objValue = objValue.Elem()
	} else {
		return errors.New("result must be addressable (a pointer)")
	}

	for i := 0; i < objType.NumField(); i++ {
		field = objType.Field(i)
		fieldValue = objValue.Field(i)

		fieldName = field.Name
		if unicode.IsLower(rune(fieldName[0])) {
			continue
		}
		og := m[field.Tag.Get("tf")]
		if og == nil {
			continue
		}
		e := fillField(og, &fieldValue)
		if e != nil {
			return e
		}
	}

	return nil

}

func fillField(data interface{}, field *reflect.Value) error {
	switch field.Type().String() {
	case "*string":
		a, ok := data.(string)
		if !ok {
			return nil
		}
		new := reflect.ValueOf(&a)
		field.Set(new)
	case "*int":
		a, ok := data.(int)
		if !ok {
			i, e := strconv.Atoi(data.(string))
			if e != nil {
				return errors.New("Type Convert error")
			}
			a = i
		}
		new := reflect.ValueOf(&a)
		field.Set(new)
	case "string":
		a, ok := data.(string)
		if !ok {
			return nil
		}
		new := reflect.ValueOf(a)
		field.Set(new)
	case "int":
		a, ok := data.(int)
		if !ok {
			i, e := strconv.Atoi(data.(string))
			if e != nil {
				return errors.New("Type Convert error")
			}
			a = i

		}
		new := reflect.ValueOf(a)
		field.Set(new)
		//case "float64":
		//	a,ok:= data.(float64)
		//	if !ok {
		//		i, e := strconv.Atoi(data.(string))
		//		if e != nil {
		//			return errors.New("Type Convert error")
		//		}
		//		a = i
		//
		//	}
		//	new := reflect.ValueOf(a)
		//	field.Set(new)
		//case "*float":

	}

	return nil
}
