package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	engine := gin.Default()

	//gin获取get参数方式集合
	engine.GET("/info/:param2", func(c *gin.Context) {

		param1 := &UserRe{}
		c.ShouldBindQuery(param1)
		c.Writer.Write([]byte("info: "+param1.Name+"  "+param1.Pwd))

		param2 := c.Param("param2")
		fmt.Println(param2)

		param3 := c.Query("param3")
		fmt.Println(param3)

		param4 := c.DefaultQuery("param4", "wu")
		fmt.Println(param4)

		param5, exits := c.GetQuery("param5")
		if exits{
			fmt.Println(param5)
		}
		
	})
	
	//gin获取post(form-data)参数方式集合
	engine.POST("/regin", func(c *gin.Context) {

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
	})

	//gin获取json参数方式集合
	engine.POST("/jsont", func(c *gin.Context) {

		param1 := &Registjson{}
		if err := c.BindJSON(param1);err != nil{
			log.Fatal(err.Error())
		}
		fmt.Println(param1.Phone, param1.Uid)

		//返回byte切片
		c.Writer.Write([]byte("byte ok!\n"))
		//返回string
		c.Writer.WriteString("string ok!\n")
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
	})
	
	
	engine.Run(":8090")
}

type UserRe struct{
	Name string `form:"uid"`
	Pwd string `form:"pwd"`
}

type Register struct{
	Uid	string `form:"name"`
	Phone string `form:"phone"`
}

type Registjson struct{
	Uid	string `form:"name"`
	Phone string `form:"phone"`
}
type Respose struct{
	Code int
	Message string
	Data interface{}
}