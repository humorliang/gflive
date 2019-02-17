package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

//注册handler
func RegisterHandlers() *httprouter.Router {
	//新的路由
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.NotFound = http.HandlerFunc(NotFound)
	//注册
	return router
}
func main() {
	//listen->handler->validation{1.request,2.user}-.business logic->response

	//http server运行  httprouter.Router实现的server中的Handler接口ServerHTTP()方法
	r := RegisterHandlers()
	fmt.Println("********* listen :8080 **********")
	http.ListenAndServe(":8000", r)

}
