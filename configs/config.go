package configs

import (
	"os"
	"reflect"
	"strconv"
)

func setValue(f reflect.Value, value string) {
	kind := f.Kind()

	if f.IsValid() && f.CanSet() {
		if kind == reflect.String {
			f.SetString(value)
		} else if kind == reflect.Bool {
			pValue, err := strconv.ParseBool(value)
			if err == nil {
				f.SetBool(pValue)
			}
		} else if kind == reflect.Struct {
			// TODO: add strut code
		} else if kind == reflect.Array {
			// TODO: add parsing code
		} else if kind == reflect.Int64 {
			pValue, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				if !f.OverflowInt(pValue) {
					f.SetInt(pValue)
				}
			}
		} else {
			// TODO: throw error of un supported config type.
		}
	}
}

func getFromEnvVariables(config interface{}) {
	typ := reflect.TypeOf(config)

	// handle pointer here
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)

		value := os.Getenv(p.Name)
		if !p.Anonymous && len(value) > 0 {
			s := reflect.ValueOf(config).Elem()

			if s.Kind() == reflect.Struct {
				f := s.FieldByName(p.Name)
				setValue(f, value)
			}
		}
	}
}

// Initialize - loading and returning the config.
func Initialize(config interface{}) {

	configValue := reflect.ValueOf(config)
	if typ := configValue.Type(); typ.Kind() != reflect.Struct {
		// TODO: throw a error
	}

	getFromEnvVariables(config)
}
