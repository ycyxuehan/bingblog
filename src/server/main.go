package main

import (
	"server/handles/user"
	"fmt"
	"net/http"
	"server/handles/category"
	"server/router"
	"server/handles/article"
	"server/connections"
	"server/config"
	"github.com/gin-gonic/gin"
)

func main(){
	var err error
	config.Conf, err = config.New("conf/app.conf")
	if err != nil {
		return
	}
	connections.MySQL.ConnectURI(config.Conf.Get("dburi"))
	defer connections.MySQL.Close()
	fmt.Println(connections.MySQL.URI)
	host := config.Conf.Get("host")
	port := config.Conf.Get("port")	
	if port == "" {
		port = "8080"
	}
	engine := gin.Default()
	
	//registe route
	router.RegistRoute(engine, article.RouterMap...)
	router.RegistRoute(engine, category.RouterMap...)
	router.RegistRoute(engine, user.RouterMap...)
	//
	fmt.Println("serving on ",fmt.Sprintf("%s:%s", host, port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), engine)
}