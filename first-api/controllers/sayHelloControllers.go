package controllers

import "net/http"

type Hello struct{} 

func HandlerSayHello(w http.ResponseWriter, r *http.Request){
	var hello Hello
	_=hello

}