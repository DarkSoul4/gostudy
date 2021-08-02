/*
 * @Descripttion:
 * @version:
 * @Author: Jie
 * @Date: 2021-07-30 14:55:55
 * @LastEditTime: 2021-07-30 17:36:29
 */
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type UserRe struct {
	Name string `form:"name"`
	Pwd  string `form:"pwd"`
}

type Register struct {
	Uid   string `form:"uid"`
	Phone string `form:"phone"`
}

type Registjson struct {
	Uid   string `form:"uid"`
	Phone string `form:"phone"`
}

type Respose struct {
	Code    int
	Message string
	Data    interface{}
}

//@Summary 测试query参数获取
// @Description 各种get参数
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param  name query string true "测试.ShouldBindQuery"
// @Param  pwd query string true "测试.ShouldBindQuery"
// @Param  param2 query string true "测试.Param"
// @Param  param3 query string true "测试.Query"
// @Param  param4 query string true "测试.DefaultQuery"
// @Param  param5 query string true "测试.GetQuery"
// @Success 200 {string} string "{"msg": "ok"}"
// @Failure 500 {string} string "{"msg": "who are you"}"
// @Router /info [get]
func test_get(c *gin.Context) {

	param1 := &UserRe{}
	c.ShouldBindQuery(param1)

	param2 := c.Param("param2")
	fmt.Println(param2)

	param3 := c.Query("param3")
	fmt.Println(param3)

	param4 := c.DefaultQuery("param4", "wu")
	fmt.Println(param4)

	param5, exits := c.GetQuery("param5")
	if exits {
		fmt.Println(param5)
		resp := Respose{Code: 200, Message: "sucess", Data: "ok"}
		c.JSON(200, &resp)
	}
}

//@Summary 测试form参数获取
// @Description 各种get参数
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param  uid formData string true "测试.ShouldBind"
// @Param  phone formData string true "测试.ShouldBind"
// @Param  param2 formData string true "测试.PostForm"
// @Param  param3 formData string true "测试.DefaultPostForm"
// @Param  param4 formData string true "测试.GetPostForm"
// @Success 200 {string} string "ok"
// @Failure 500 {string} string "{"msg": "ERROR"}"
// @Router /regin [post]
func test_form(c *gin.Context) {

	param1 := &Register{}
	c.ShouldBind(param1)
	fmt.Println(param1.Phone, param1.Uid)

	param2 := c.PostForm("param2")
	fmt.Println(param2)

	param3 := c.DefaultPostForm("param3", "ooo")
	fmt.Println(param3)

	param4, exits := c.GetPostForm("param4")
	fmt.Println(param4, exits)

	c.Writer.Write([]byte("ok!"))
}

//@Summary 测试json参数获取
// @Description 各种get参数
// @Tags 测试
// @Accept json
// @Produce json
// @Param  param body main.Registjson true "测试.BindJSON"
// @Success 200 {object} main.Respose
// @Failure 500 {string} string "{"msg": "error"}"
// @Router /jsont [post]
func test_json(c *gin.Context) {
	param1 := &Registjson{}
	if err := c.BindJSON(param1); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(param1.Phone, param1.Uid)
	//返回json
	// c.JSON(
	// 	200,
	// 	map[string]interface{}{
	// 		"statue":1,
	// 		"message":"sucess",
	// 		"data":param1,
	// 	},
	// )
	//返回自定义格式
	resp := Respose{Code: 200, Message: "sucess", Data: param1}
	c.JSON(200, &resp)
}
