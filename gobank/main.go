package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func UserData(name string, age int, email string) *User {
	return &User{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func main() {
	user1 := UserData("Ridho", 22, "Ridhomp@gmail.com")
	fmt.Println(user1)
}
