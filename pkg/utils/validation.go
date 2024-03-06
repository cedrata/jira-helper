package utils

import (
	"fmt"
	"reflect"
)

func ValidateStruct[T any](data T) error {
	typeOfData := reflect.TypeOf(data)
	valueOfData := reflect.ValueOf(data)

	if kind := typeOfData.Kind(); kind != reflect.Struct {
		return fmt.Errorf("expected a struct, got instead %v", kind)
	}

	for fieldIndex := range typeOfData.NumField() {
		field := typeOfData.Field(fieldIndex)
		fieldValue := valueOfData.Field(fieldIndex)

		tagValue, ok := field.Tag.Lookup("validation")
		if !ok {
			continue
		}

		isValueZero := reflect.DeepEqual(
			fieldValue.Interface(),
			reflect.Zero(field.Type).Interface(),
		)
		if tagValue == "required" && isValueZero {
			return fmt.Errorf("field %s is required but not set", field.Name)
		}
	}

	return nil
}
