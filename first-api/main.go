package main

import (
	"flag"
	"net/http"

	"github.com/cocoaster/first-api/controllers"
)


func main(){
	listenAddrs:= flag.String("listenaddrs", ":8090", "t1odo")
	flag.Parse()
	http.HandleFunc("/sayHello", controllers.HandlerSayHello)


	http.ListenAndServe(*listenAddrs, nil)
}