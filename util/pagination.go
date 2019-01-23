package util

import (
	"github.com/gin-gonic/gin",
	"go-gin-server/common"
)

func GetPage(c *gin.Context) {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * common.PageSize
	}
	return result
}
