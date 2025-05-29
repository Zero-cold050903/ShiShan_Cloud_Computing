package api

import (
	"fmt"
	"math/rand"
	"os"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/config"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/types"
	"github.com/gin-gonic/gin"
)
var router = gin.Default()
var upload = router.Group("/upload")
var download = router.Group("/download")
var file *gin.FileHeader
var err error
func distributeID() string {
	// 生成ID
	rand.Rand()		// 随机数种子
	id := rand.Intn(10000000000000)
	return string(id) 
}
func (q types.TaskQueue)AddTask() {
	// 添加任务
	q.Queue[TaskNum] = task{
		TaskID:     distributeID(),
		TaskName:   fmt.Fscanln(),
		TaskType:   ,
		TaskStatus: config.WAITING,
		TaskFile: 
	}
}
func OpenULapi() {
	// 上传文件
	upload.POST("/document", uploadDocument)
	upload.POST("/music", uploadMusic)
	upload.POST("/file", uploadFile)
	upload.POST("/image", uploadImage)
	router.Run(":8080")
}
func CloseULapi() {
	router = nil
	upload = nil
	file = nil
	err = nil
}
func OpenDLapi() {
	//
	download.GET("/document",downloadDocument)
	download.GET("/music",downloadMusic)
	download.GET("/file",downloadFile)
	download.GET("/image",downloadImage)
	router.Run(":8080")
}
func CloseDLapi() {
	router = nil
	download = nil
	file = nil
	err = nil
}
func uploadDocument(c *gin.Context) error {
	// 上传文件
	file, err := c.FormFile("file")
	if err != nil {
		routes.ErrorResponse(c, 400, "上传文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = os.Mkdir("./USERFILE", 0777)
	if err != nil {
		routes.ErrorResponse(c, 500, "创建文件夹失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = c.SaveUploadedFile(file, "./USERFILE/"+file.Filename)
	if err != nil {
		routes.ErrorResponse(c, 500, "保存文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	routes.SuccessResponse(c, 200, "上传文件成功", gin.H{
		"filename": file.Filename,
	})
	return nil
}

func uploadMusic(c *gin.Context) error {
	// 上传文件
	file, err := c.FormFile("file")
	if err!= nil {
		routes.ErrorResponse(c, 400, "上传文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = os.Mkdir("./USERFILE", 0777)
	if err!= nil {
		routes.ErrorResponse(c, 500, "创建文件夹失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = c.SaveUploadedFile(file, "./USERFILE/"+file.Filename)
	if err!= nil {
		routes.ErrorResponse(c, 500, "保存文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}

	}
	routes.SuccessResponse(c, 200, "上传文件成功", gin.H{
		"filename": file.Filename,
	})
	return nil
}
func uploadFile(c *gin.Context) error {
	// 上传文件
	file, err : = c.FormFile("file")
	if err != nil {
		routes.ErrorResponse(c,400, "上传文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = os.Mkdir("./USERFILE",0777)
	if err != nil {
		routes.ErrorResponse(c,500, "创建文件夹失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: "Create folder failed",
		}
	}
	err = c.SaveUploadedFile(file, "./USERFILE/"+file.Filename)
	if err!= nil {
		routes.ErrorResponse(c,500, "保存文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: "Save file failed",
		}
	}
	routes.SuccessResponse(c,200, "上传文件成功", gin.H{
		"filename": file.Filename,
	})
	return nil
}
func uploadVideo(c *gin.Context) error {
	file, err := c.FormFile("file")
	if err!= nil {
		routes.ErrorResponse(c, 400, "上传文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = os.Mkdir("./USERFILE", 0777)
	if err!= nil {
		routes.ErrorResponse(c, 500, "创建文件夹失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}		
	}
	err = c.SaveUploadedFile(file, "./USERFILE/"+file.Filename)
	if err!= nil {
		routes.ErrorResponse(c, 500, "保存文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	routes.SuccessResponse(c, 200, "上传文件成功", gin.H{
		"filename": file.Filename,
	})
}
func uploadPicture(c *gin.Context) error {
	file, err := c.FormFile("file")
	if err!= nil {
		routes.ErrorResponse(c, 400, "上传文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = os.Mkdir("./USERFILE", 0777)
	if err!= nil {
		routes.ErrorResponse(c, 500, "创建文件夹失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	err = c.SaveUploadedFile(file, "./USERFILE/"+file.Filename)
	if err!= nil {
		routes.ErrorResponse(c, 500, "保存文件失败")
		return gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		}
	}
	routes.SuccessResponse(c, 200, "上传文件成功", gin.H{
		"filename": file.Filename,
	})
	return nil
}
func downloadDocument(c *gin.Context) error {
	
}
func downloadMusic(c *gin.Context) {

}
func downloadFile(c *gin.Context) {

}
func downloadVideo(c *gin.Context) {

}
func downloadPicture(c *gin.Context) {

}
