package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateInputs(dataSet interface{}) (bool, map[string]string) {

	// var validate * validator.Validate
	validate := validator.New()

	err := validate.Struct(dataSet)

	if err != nil {

		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		errors := make(map[string]string)

		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors){

			field, _ :=reflected.Type().FieldByName(err.StructField())
			var name string

			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			// fmt.Println(err.Tag(), err.Type(), err.StructField())

			switch err.Tag(){


			case "required":
				errors[name] = "The "+ name + " is required"
				break
			case "email":
				errors[name] = "The "+ name + " should be a valid email"
				break
			case "eqfiled":
				errors[name] = "The "+ name + " should be equal to the "+ err.Param()
				break
			case "default":
				errors[name] = "The "+ name + " is invalid"
				break

			}

		}
		return false, errors

		}
	return true, nil
}