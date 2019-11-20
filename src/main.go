package main

import (
	"net/http"

	"./controllers"
	"./env"
	"github.com/astaxie/beego/logs"
)

func main() {
	env.InitEnv()
	mux := http.NewServeMux()
	mux.HandleFunc("/go-helloWorld", controllers.GoHelloWorld)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/logout", controllers.Logout)
	mux.HandleFunc("/register", controllers.Register)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		logs.Error(err)
	}
}
