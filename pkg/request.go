package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBindWith(request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "parse Request Error",
			"error": err.Error(),
		})
		log.Println("ParseRequest Result", request)
		log.Println("ParseRequest Error", err.Error())
		return err
	}
	return nil
}

func SuccessResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func CheckErr(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
}
