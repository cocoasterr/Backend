package helper

import (
	"encoding/json"
	"net/http"
)


func ResponseJSON(w http.ResponseWriter, code int, payload interface{}){
	response,_ := json.Marshal(payload)
	w.Header().Add("content-type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetJson(input interface{},w http.ResponseWriter, r *http.Request)interface{}{
	var response map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input);err != nil{
		response = map[string]string{"message": "Decode json failed!"}
		ResponseJSON(w, http.StatusBadRequest, response)
		return nil
	}
	return input
}

