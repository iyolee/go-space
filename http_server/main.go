package main

import (
	"fmt"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello"
	w.Write([]byte(str))
}

func main() {
	s := "hello 李"
	for i, r := range s {
		fmt.Println(i, r)
	}
	http.HandleFunc("/post/go", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
