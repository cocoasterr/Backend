package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct{
	listenAddrs string
}

func NewAPiServer(listenAddrs string) *APIServer{
	return &APIServer{
		listenAddrs: listenAddrs,
	}
}

func (s *APIServer) Run(){
	router:= mux.NewRouter()
	router.HandleFunc("/account", s.HandleAccount)
	
}

func (s *APIServer) HandleAccount (w http.ResponseWriter,r *http.Request) error{
	return nil
}
func (s *APIServer) HandleGetAccount (w http.ResponseWriter,r *http.Request) error{
	return nil
}
func (s *APIServer) HandleCreateAccount (w http.ResponseWriter,r *http.Request) error{
	return nil
}
func (s *APIServer) HandleDeleteAccount (w http.ResponseWriter,r *http.Request) error{
	return nil
}
func (s *APIServer) HandleTransferAccount (w http.ResponseWriter,r *http.Request) error{
	return nil
}
