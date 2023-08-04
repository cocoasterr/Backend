package main

import "fmt"

type user struct{
	name string
	email string
	age int
}

func (u user)getEmail() string{
	return	u.email
}

type server struct{
	maxConn int
	id string
	tls bool
}

func newServer(maxconn int, id string, tls bool) server{
	return server{
		maxConn: maxconn,
		id: id,
		tls: tls,
	}
}

func main(){
	// user1 := user{
	// 	name : "ridho",
	// 	email: "blabla@gmail.com",
	// }
	// fmt.Println(user1.getEmail())
	showServer := newServer(5,"wqeqweqwe76", true)
	fmt.Println(showServer.id)


}