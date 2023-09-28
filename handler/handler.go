package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
