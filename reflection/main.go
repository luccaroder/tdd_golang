package main

import "reflect"

func Iteration(x interface{}, f func(input string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			f(field.String())
		case reflect.Struct:
			Iteration(field.Interface(), f)
		}

	}
}
