package main

import "net/http"

func main(){

	http.HandleFunc("/user/homeholder", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":10000",nil)
}
