package main

import (
	"log"
	"net/http"

	"github.com/SiriusLLL/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 建立路由
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	// 将路由r与根路径"/"进行绑定。
	// 意味着当用户访问根路径时，将使用这个路由来处理请求。
	http.Handle("/", r)
	// 启动监听地址"localhost:9010"
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
