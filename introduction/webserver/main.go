package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
