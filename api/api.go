package api

import (
	"fuck-online-course/model"
	"fuck-online-course/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var RequestLimit = util.NewRequestLimitService(10*time.Second, 5)

func GetAnswers(c *gin.Context) {

	if RequestLimit.IsAvailable() {
		RequestLimit.Increase()
		goto Service
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求太过频繁"})
		return
	}

Service:
	name := c.Query("name")

	var course model.Courses
	rows, count, err := course.FindQuestion(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": rows,
		"total": count,
	})
}
