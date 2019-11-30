package main

import (
	"net/http"

	"./controllers"
	"./env"
	"./models"
	"github.com/astaxie/beego/logs"
)

func main() {
	env.InitEnv()
	res, _ := models.GetFavourGoods(2,"")
	logs.Info(res)
	mux := http.NewServeMux()
	mux.HandleFunc("/go-helloWorld", controllers.GoHelloWorld)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/logout", controllers.Logout)
	mux.HandleFunc("/register", controllers.Register)
	// 收藏夹
	mux.HandleFunc("/getFolder", controllers.GetFolder)
	mux.HandleFunc("/getFavourGoods", controllers.GetFavourGoods)
	// 货品种类
	mux.HandleFunc("/getGoodsType", controllers.GetGoodsType)
	// 主页
	mux.HandleFunc("/getGoodsByType", models.GetGoodsByType)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		logs.Error(err)
	}
}
