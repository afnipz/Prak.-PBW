package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case"POST":
			w.Write([]byte("post"))
		case "GET":
			w.Write([]byte("get"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("Server started at localhost:1190")
	http.ListenAndServe(":1190", nil)
}