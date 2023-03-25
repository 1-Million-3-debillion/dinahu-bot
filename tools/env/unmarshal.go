package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Unmarshal(config interface{}) error {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := os.Getenv(field.Tag.Get("env"))
		f := v.FieldByName(field.Name)
		if f.Kind() == reflect.Struct {
			err := Unmarshal(f.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		if value != "" {
			switch f.Kind() {
			case reflect.String:
				f.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				iv, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to parse %s as int: %v", field.Name, err)
				}
				f.SetInt(iv)
			case reflect.Bool:
				bv, err := strconv.ParseBool(value)
				if err != nil {
					return fmt.Errorf("failed to parse %s as bool: %v", field.Name, err)
				}
				f.SetBool(bv)
			case reflect.Float32, reflect.Float64:
				fv, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return fmt.Errorf("failed to parse %s as float: %v", field.Name, err)
				}
				f.SetFloat(fv)
			default:
				return errors.New("invalid type")
			}
		}
	}
	return nil
}
