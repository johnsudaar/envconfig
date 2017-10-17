package envconfig

import (
	"os"
	"reflect"
)

func Build(a interface{}) {
	t := reflect.ValueOf(a).Elem()
	for i := 0; i < t.Type().NumField(); i++ {
		field := t.Type().Field(i)
		variable := field.Tag.Get("env")
		if variable == "" {
			continue
		}
		value := os.Getenv(variable)
		if value == "" {
			continue
		}
		t.FieldByName(field.Name).SetString(value)
	}
}
