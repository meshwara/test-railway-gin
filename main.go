package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/file/:tipe", func(ctx *gin.Context) {
		tipe := ctx.Param("tipe")
		if tipe == "success" {
			var path = "public/berkas.pdf"
			ctx.Header("Content-Description", "File Transfer")
			ctx.Header("Content-Transfer-Encoding", "binary")
			ctx.Header("Content-Disposition", "attachment; filename=berkas.pdf")
			ctx.Header("Content-Type", "application/octet-stream")
			ctx.File(path)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Data tidak valid",
			})
		}
	})
	r.Run()
}
