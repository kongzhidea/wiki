package mreflections

import (
	"errors"
	"fmt"
	"reflect"
)


// 只处理int,int8~int64类型。
func GetFieldInt64Value(obj interface{}, fieldName string) (int64, error) {
	t := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return 0, errors.New("parse field value error: obj is not struct")
	}
	field := t.FieldByName(fieldName)

	if field.Kind() == reflect.Invalid {
		return 0, errors.New("parse field value error: not found field " + fieldName)
	}
	// 如果是指针，获取指向的结构
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	if field.Kind() == reflect.Int || field.Kind() == reflect.Int8 || field.Kind() == reflect.Int16 || field.Kind() == reflect.Int32 || field.Kind() == reflect.Int64 {
		return field.Int(), nil
	}

	return 0, errors.New(fmt.Sprintf("parse %v field value error, field type not support,%v is %v", "int", fieldName, field.Kind().String()))
}

// 只处理string类型
func GetFieldStringValue(obj interface{}, fieldName string) (string, error) {
	t := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", errors.New("parse field value error: obj is not struct")
	}
	field := t.FieldByName(fieldName)

	if field.Kind() == reflect.Invalid {
		return "", errors.New("parse field value error: not found field " + fieldName)
	}
	// 如果是指针，获取指向的结构
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	if field.Kind() == reflect.String {
		return field.String(), nil
	}

	return "", errors.New(fmt.Sprintf("parse %v field value error, field type not support,%v is %v", "string", fieldName, field.Kind().String()))
}

// 只处理bool类型
func GetFieldBoolValue(obj interface{}, fieldName string) (bool, error) {
	t := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return false, errors.New("parse field value error: obj is not struct")
	}
	field := t.FieldByName(fieldName)

	if field.Kind() == reflect.Invalid {
		return false, errors.New("parse field value error: not found field " + fieldName)
	}
	// 如果是指针，获取指向的结构
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	if field.Kind() == reflect.Bool {
		return field.Bool(), nil
	}

	return false, errors.New(fmt.Sprintf("parse %v field value error, field type not support,%v is %v", "bool", fieldName, field.Kind().String()))
}
