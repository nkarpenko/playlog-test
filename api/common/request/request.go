package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// HandlerFunc used as helper proxy function
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// ParseFormDataOrJSONBody parses the request data
// from From value or JSON body into the desired struct
func ParseFormDataOrJSONBody(r *http.Request, o interface{}) error {
	// Get content type
	contentType := r.Header.Get("Content-type")
	if contentType != "" {
		contentType = strings.Split(contentType, ";")[0]
	}

	switch contentType {
	case "application/x-www-form-urlencoded":
		r.ParseForm()
		return ParseFormData(r.Form, o)
	case "multipart/form-data":
		r.ParseMultipartForm(512000) // 512 Kb
		return ParseFormData(r.Form, o)
	case "application/json":
		return ParseJSONBody(r.Body, o)
	default:
		return errors.New("Unsupported content-type")
	}
}

// ParseFormData into the desired struct
// Unexpected data are skipped
func ParseFormData(form url.Values, o interface{}) error {
	ot := reflect.TypeOf(o)
	ov := reflect.ValueOf(o)

	var typ reflect.Type
	var val reflect.Value
	if ot.Kind() == reflect.Ptr {
		typ = ot.Elem()
		val = ov.Elem()
	} else {
		typ = ot
		val = ov
	}

	if typ.Kind() != reflect.Struct {
		return errors.New("ParseFormData - cannot update a non struct object")
	}

	if val.CanSet() == false {
		return errors.New("ParseFormData - the output struct must be passed as a reference")
	}

	// Flatten the from data
	data := map[string]string{}
	for name, values := range form {
		data[name] = values[0]
	}

	// Try to update the struct fields
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		f := val.Field(i)
		tag := sf.Tag.Get("json")
		value, ok := data[tag]
		if ok && f.CanSet() {
			switch sf.Type.Kind() {
			case reflect.String:
				f.SetString(value)
			case reflect.Int:
				v, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				f.SetInt(v)
			case reflect.Float32:
			case reflect.Float64:
				v, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				f.SetFloat(v)
			case reflect.Bool:
				v, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				f.SetBool(v)
				// TODO: Handle other non-string type conversion
			}
		}
	}
	return nil
}

// ParseJSONFormData from form values
func ParseJSONFormData(form url.Values, object interface{}) error {
	fmt.Printf("%s\n", form)

	var jsonData string
	for v := range form {
		jsonData = v
		break
	}
	fmt.Printf("%s\n", jsonData)
	return json.Unmarshal([]byte(jsonData), &object)
}

// ParseJSONBody parses the request JSON body
// into the desired struct
func ParseJSONBody(body io.ReadCloser, o interface{}) error {
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(raw), &o)
}

// GetVarInt from the gorilla mux vars
func GetVarInt(vars map[string]string, key string) (int, error) {
	raw, ok := vars[key]
	if ok == false {
		return 0, fmt.Errorf("%s not found in vars", key)
	}
	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %s to int param from vars:%+v", key, err)
	}
	return val, nil
}

// GetVarInt64 from the gorilla mux vars
func GetVarInt64(vars map[string]string, key string) (int64, error) {
	raw, ok := vars[key]
	if ok == false {
		return 0, fmt.Errorf("%s not found in vars", key)
	}
	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %s to int64 param from vars:%+v", key, err)
	}
	return val, nil
}

// GetVarString from the gorilla mux vars
func GetVarString(vars map[string]string, key string) (string, error) {
	val, ok := vars[key]
	if ok == false {
		return "", fmt.Errorf("%s not found in vars", key)
	}
	return val, nil
}
