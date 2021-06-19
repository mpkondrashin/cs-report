package main

import (
	"fmt"
	"reflect"
)

func StructCSV(structPtr interface{}) (name []string, value []string) {
	return structCSV(structPtr, "")
}

func structCSV(structPtr interface{}, prefix string) (name []string, value []string) {
	name = make([]string, 0)
	value = make([]string, 0)
	//structPtrValue := reflect.ValueOf(structPtr)
	//structValue := reflect.Indirect(structPtrValue)
	structValue := reflect.ValueOf(structPtr)
	if structValue.Kind() == reflect.Ptr {
		structValue = reflect.Indirect(structValue)
	}
	//structValue := reflect.ValueOf(data)
	structValueType := structValue.Type()
	//fmt.Println("structCSV", structValue, structValueType)
	for i := 0; i < structValue.NumField(); i++ {
		fieldType := structValueType.Field(i)
		fieldValue := structValue.Field(i)
		fieldName := fieldType.Name
		if prefix != "" {
			fieldName = fmt.Sprintf("%s.%s", prefix, fieldType.Name)
		}
		if fieldValue.Type().Kind() == reflect.Struct {
			subName, subValue := structCSV(fieldValue.Interface(), prefix)
			name = append(name, subName...)
			value = append(value, subValue...)
		} else {
			name = append(name, fieldName)
			//	fmt.Println("name = ", name)
			value = append(value, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}
	return
}

type S struct {
	A string
	B struct {
		X int
		Y int
	}
}

func main() {
	var s S
	s.A = "AD"
	s.B.X = 2
	s.B.Y = 3
	name, value := StructCSV(&s)
	fmt.Println(name)
	fmt.Println(value)
}
