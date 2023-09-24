package main

import (
	"fmt"

	helper "github.com/cocoasterr/learn-go/helper"
	model "github.com/cocoasterr/learn-go/models"
)

func main() {
	person := model.NewPerson("ridho", "ridhompra", "password", "lala@mail.com")
	res := helper.StructToMap(person)
	fmt.Println(res)
}
