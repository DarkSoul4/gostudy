/*
 * @Descripttion:
 * @version:
 * @Author: Jie
 * @Date: 2021-07-16 10:50:16
 * @LastEditTime: 2021-07-30 17:36:20
 */
package main

import (
	_ "gin_helloworld/docs"

	"github.com/gin-gonic/gin"
)

var pre = "api/v1"
var swagHandler gin.HandlerFunc

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8090
// @BasePath /api/v1
func main() {
	// 创建一个默认的路由引擎
	engine := gin.Default()

	//gin获取get参数方式集合
	engine.GET(pre+"/info", test_get)

	//gin获取post(form-data)参数方式集合
	engine.POST(pre+"/regin", test_form)

	//gin获取json参数方式集合
	engine.POST(pre+"/jsont", test_json)

	if swagHandler != nil {
		engine.GET("/swagger/*any", swagHandler)
	}

	engine.Run(":8090")
}
