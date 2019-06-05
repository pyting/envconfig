package envconfig

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	typeError = "expected a pointer to a struct"
)

func Parse(prefix string, spec interface{}) error {

	rv := reflect.ValueOf(spec)
	if rv.Kind() != reflect.Ptr {
		return errors.New(typeError)
	}

	re := rv.Elem()
	if re.Kind() != reflect.Struct {
		return errors.New(typeError)
	}

	for i := 0; i < re.NumField(); i++ {
		var sv interface{}
		fieldInfo := re.Type().Field(i)

		defV := fieldInfo.Tag.Get("default")
		v := os.Getenv(strings.ToUpper(prefix) + "_" + fieldInfo.Tag.Get("env"))
		switch fieldInfo.Type.Kind() {
		case reflect.String:
			if v == "" {
				sv = defV
			} else {
				sv = v
			}

		case reflect.Bool:
			var t bool
			var err error
			if v == "" {
				if defV != "" && defV != "false" {
					t, err = strconv.ParseBool(defV)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseBool(v)
				if err != nil {
					return err
				}
			}
			sv = t
		case reflect.Int:
			var t int64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseInt(defV, 10, 32)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseInt(v, 10, 32)
				if err != nil {
					return err
				}
			}
			sv = int(t)
		case reflect.Int8:
			var t int64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseInt(defV, 10, 8)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseInt(v, 10, 8)
				if err != nil {
					return err
				}
			}
			sv = int8(t)
		case reflect.Int16:
			var t int64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseInt(defV, 10, 16)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseInt(v, 10, 16)
				if err != nil {
					return err
				}
			}
			sv = int16(t)
		case reflect.Int32:
			var t int64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseInt(defV, 10, 32)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseInt(v, 10, 32)
				if err != nil {
					return err
				}
			}
			sv = int32(t)
		case reflect.Int64:
			var t int64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseInt(defV, 10, 64)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
			}
			sv = t
		case reflect.Uint:
			var t uint64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseUint(defV, 10, 32)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseUint(v, 10, 32)
				if err != nil {
					return err
				}
			}
			sv = uint(t)
		case reflect.Uint8:
			var t uint64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseUint(defV, 10, 8)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseUint(v, 10, 8)
				if err != nil {
					return err
				}
			}
			sv = uint8(t)
		case reflect.Uint16:
			var t uint64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseUint(defV, 10, 16)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseUint(v, 10, 16)
				if err != nil {
					return err
				}
			}
			sv = uint16(t)
		case reflect.Uint32:
			var t uint64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseUint(defV, 10, 32)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseUint(v, 10, 32)
				if err != nil {
					return err
				}
			}
			sv = uint32(t)
		case reflect.Uint64:
			var t uint64
			var err error
			if v == "" {
				if defV != "" && defV != "0" {
					t, err = strconv.ParseUint(defV, 10, 64)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseUint(v, 10, 64)
				if err != nil {
					return err
				}
			}
			sv = t
		case reflect.Float32:
			var t float64
			var err error
			if v == "" {
				if defV != "" {
					t, err = strconv.ParseFloat(defV, 32)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
			}
			sv = float32(t)
		case reflect.Float64:
			var t float64
			var err error
			if v == "" {
				if defV != "" {
					t, err = strconv.ParseFloat(defV, 64)
					if err != nil {
						return err
					}
				}
			} else {
				t, err = strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
			}
			sv = t
		}
		re.Field(i).Set(reflect.ValueOf(sv))
	}
	return nil
}
