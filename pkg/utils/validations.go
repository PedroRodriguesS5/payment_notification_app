package tools

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

const tagCustom = "errormgs"

func ErrorTagFunc[T interface{}](obj interface{}, snp string, fieldName, actualTag string) error {
	o := obj.(T)

	if !strings.Contains(snp, fieldName) {
		return nil
	}

	fieldArr := strings.Split(snp, ".")
	rsf := reflect.TypeOf(o)

	for i := 0; i < len(fieldArr); i++ {
		field, found := rsf.FieldByName(fieldArr[i])
		if found {
			if fieldArr[i] == fieldName {
				customMessage := field.Tag.Get(tagCustom)
				if customMessage != "" {
					return fmt.Errorf("%s: %s (%s)", fieldName, customMessage, actualTag)
				}
				return nil
			} else {
				if field.Type.Kind() == reflect.Ptr {
					rsf = field.Type.Elem()
				} else {
					rsf = field.Type
				}
			}
		}
	}
	return nil
}

func ValidateFunc[T interface{}](obj interface{}, validate *validator.Validate) (errs error) {
	o := obj.(T)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in validate: ", r)
			errs = fmt.Errorf("can't validte %+v", r)
		}
	}()

	if err := validate.Struct(o); err != nil {
		errorValid := err.(validator.ValidationErrors)
		for _, e := range errorValid {
			snp := e.StructNamespace()
			errmgs := ErrorTagFunc[T](obj, snp, e.Field(), e.ActualTag())
			if errmgs != nil {
				errs = errors.Join(errs, fmt.Errorf("%w", errmgs))
			} else {
				errs = errors.Join(errs, fmt.Errorf("%w", e))
			}
		}
	}
	if errs != nil {
		return errs
	}

	return nil
}
