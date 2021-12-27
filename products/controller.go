package products

import "github.com/gin-gonic/gin"

func Post(context *gin.Context) {
	var service Service
	service.load()
	var response CartResponse
	if err := context.BindJSON(&response); err != nil {
		return
	}
}
