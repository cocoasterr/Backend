package main

import (
	"fmt"
	"reflect"
)


type User struct{
	Id int
	Name string
	Address string
	Age int8
}


func mappingStruct(data any) map[string]any{
	res:= make(map[string]any)
	value := reflect.ValueOf(data)
	typ:= value.Type()

	for i := 0; i < value.NumField(); i++ {
		fieldKey:= typ.Field(i).Name
		fieldValue:= value.Field(i).Interface()
		res[fieldKey] = fieldValue
	}
	return res
}

func main(){
	ridho:= User{Id:1, Name: "Ridho", Address: "Indonesia", Age: 20}
	res := mappingStruct(ridho)
	fmt.Println(res)
}