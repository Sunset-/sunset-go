package validates

import (
	"fmt"
	"reflect"
)

func ValidateStruct(model interface{}) {
	modelType := getElem(model)
	if modelType.Kind() != reflect.Struct {
		return
	}
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Struct {
		modelValue = modelValue.Elem()
	}

	fNum := modelType.NumField()
	for i := 0; i < fNum; i++ {

		field := modelType.Field(i)
		validateTag := field.Tag.Get("validate")
		validateField(model, field, modelValue.FieldByName(field.Name), validateTag)
	}
}

func getElem(inter interface{}) reflect.Type {
	modelType := reflect.TypeOf(inter)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	return modelType
}

func validateField(model interface{}, field reflect.StructField, value reflect.Value, validateTag string) {
	fmt.Println(field.Name, value.String())
}
