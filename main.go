package main

import (
	"fuck-online-course/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/get_answer", api.GetAnswers)
	err := r.Run(":8089") // listen and serve on 0.0.0.0:8089
	if err != nil {
		panic("start failed")
	}
}
