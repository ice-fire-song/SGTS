package controllers

import "net/http"

func GoHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("go hello world"))
}
