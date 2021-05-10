//Написать функцию, которая принимает на вход структуру in (struct или кастомную struct)
//и values map[string]interface{} (key - название поля структуры, которому нужно присвоить
//value этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с помощью
//пакета reflect. Функция может возвращать только ошибку error. Написать к данной функции тесты
//(чем больше, тем лучше - зачтется в плюс).
package main

import (
	"errors"
	"fmt"
	"reflect"
)
type User struct {
	Username string `json:"user"`
	Password string
	Number int
}

func main() {

	in := &User{
		Username: "Yulia",
		Password: "080808",
		Number: 888,
	}

	myMap := map[string]interface{} {
		"Username": "J",
		"Password": "070707",
		"Number": 777,
	}

	fmt.Printf("original struct: %v\n", in)

	err := PrintStruct(in, myMap)
	if err !=nil {
		err.Error()
	}

	fmt.Printf("changed struct: %v\n", in)
}


func PrintStruct(in interface{}, myMap map[string]interface{}) error {
	if in == nil {
		return errors.New("entry is nil")
	}

	inType := reflect.ValueOf(in)
	if inType.Kind() == reflect.Ptr {
		inType = inType.Elem()
	}

	if inType.Type().Kind() != reflect.Struct {
		return fmt.Errorf("is not struct; kind %s", inType.Type().Kind())
	}

	for key, ok := range myMap {
		field := inType.FieldByName(key)
		tp := reflect.ValueOf(ok)
		if field.Type().Kind() != tp.Type().Kind() {
			return errors.New("type field is not correct")
		}
		field.Set(tp)

	}
	return nil
}
