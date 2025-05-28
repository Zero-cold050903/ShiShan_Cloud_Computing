package routes

import (
	"github.com/gin-gonic/gin"
)

// Init the root router
var Router = gin.Default()

func openUPLOAD() {
	upload := Router.Group("/upload")
	upload.POST("/picture", uploadPicture)
	upload.POST("/video", uploadVideo)
	upload.POST("/file", uploadFile)
	upload.POST("/music", uploadMusic)
	upload.POST("/document", uploadDocument)
}

func openDownLoad() {
	DownLoad := Router.Group("/download")
	DownLoad.GET("/picture", downloadPicture)
	DownLoad.GET("/video", downloadVideo)
	DownLoad.GET("/file", downloadFile)
	DownLoad.GET("/music", downloadMusic)
	DownLoad.GET("/document", downloadDocument)
}
