package main

import "net/http"

func myHandler(reponse http.ResponseWriter, request *http.Request) {
	reponse.Write([]byte("hello, word"))
}
func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe(":8080", nil)
}
