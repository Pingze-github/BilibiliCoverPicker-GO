package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func server() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例

	//添加中间件
	router.Use(Middleware)
	//注册接口
	router.GET("/simple/server/get", GetHandler)
	router.POST("/simple/server/post", PostHandler)
	router.PUT("/simple/server/put", PutHandler)
	router.DELETE("/simple/server/delete", DeleteHandler)
	//监听端口
	http.ListenAndServe(":8005", router)
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}

func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}

func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}

func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}

func main() {
	server()
}