package unit_tes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Lenght[T any](params T) T{
	fmt.Println(params)
	fmt.Println(reflect.TypeOf(params))
	return params
}

func MultipleTypeGeneric[T1 any, T2 any](param1 T1, param2 T2){
	fmt.Println(param1)
	fmt.Println(param2)
}

func IsComparable[T comparable](value1, value2 T)bool{
	return value1 == value2
}

func TestSample(t *testing.T){
	inputStr := "ridho"
	resultString := Lenght[string](inputStr)
	fmt.Println(resultString)
	assert.Equal(t,inputStr, resultString)
	
	inputInt := 8080
	resultInt := Lenght[int](inputInt)
	fmt.Println(resultInt)
	assert.Equal(t,inputInt, resultInt)
}
func TestMultiple(t *testing.T){
	MultipleTypeGeneric[string, int]("Ridho", 1080)
}


func TestCompare(t *testing.T) {
	res:= IsComparable[string]("ridho", "ridho")
	fmt.Println(res)
}