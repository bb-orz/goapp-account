package restful

import (
	"errors"
	"fmt"
	"github.com/bb-orz/goinfras/XGin"
	"github.com/bb-orz/goinfras/XOss/XQiniuOss"
	"github.com/gin-gonic/gin"
	"goapp/common"
	"goapp/restful/middleware"
	"mime/multipart"
	"sync"
)

/*
API层，调用相关Service，封装响应返回，并记录日志
*/

func init() {
	var once sync.Once
	once.Do(func() {
		// 初始化时注册该模块API
		XGin.RegisterApi(new(ResourceApi))
	})
}

type ResourceApi struct{}

// 设置该模块的API Router
func (api *ResourceApi) SetRoutes() {
	engine := XGin.XEngine()

	engine.Static("/image", common.UploadImagesPath)
	engine.Static("/file", common.UploadFilesPath)
	engine.Static("/video", common.UploadVideosPath)
	engine.Static("/audio", common.UploadAudiosPath)

	// 上传资源相关路由
	engine.MaxMultipartMemory = 15 << 20 // 设置最大上传大小为15M
	userGroup := engine.Group("/resource", middleware.JwtAuthMiddleware())
	userGroup.GET("/get_qiniu_upload_token", api.getQiniuUploadTokenHandler)
	userGroup.POST("/upload_image", api.uploadImageHandler)
	userGroup.POST("/upload_file", api.uploadFileHandler)
	userGroup.POST("/upload_video", api.uploadVideoHandler)
	userGroup.POST("/upload_audio", api.uploadAudioHandler)

}

func (api *ResourceApi) getQiniuUploadTokenHandler(ctx *gin.Context) {

	upToken := XQiniuOss.XClient().SimpleUpload()
	if upToken == "" {
		_ = ctx.Error(common.ErrorOnNetworkRequest("Get Qiniu Oss Client Upload Token Fail"))
		return
	}

	ctx.Set(common.ResponseDataKey, common.ResponseOK(gin.H{"upToken": upToken}))
}

func (api *ResourceApi) uploadImageHandler(ctx *gin.Context) {
	var err error
	var fileHeader *multipart.FileHeader
	// var file multipart.File
	if fileHeader, err = ctx.FormFile("file"); err != nil {
		_ = ctx.Error(err)
		return
	}
	fmt.Printf("FileHeader: %#v \n", fileHeader.Header.Values("Content-Type"))

	contentType := fileHeader.Header.Values("Content-Type")[0]
	AllowImageTypes := []string{"image/jpeg", "image/png", "image/gif"}
	if !common.IsStringItemExist(AllowImageTypes, contentType) {
		_ = ctx.Error(errors.New("Content-Type Is Not Allowed! "))
		return
	}

	// Upload the file to specific dst.
	dst := common.UploadImagesPath + "/" + fileHeader.Filename
	if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {
		_ = ctx.Error(errors.New("Save Upload Images Fail "))
		return
	}

	ctx.Set(common.ResponseDataKey, nil)

}

func (api *ResourceApi) uploadFileHandler(ctx *gin.Context) {
	var err error
	var fileHeader *multipart.FileHeader
	// var file multipart.File
	if fileHeader, err = ctx.FormFile("file"); err != nil {
		_ = ctx.Error(err)
		return
	}
	fmt.Printf("FileHeader: %#v \n", fileHeader.Header.Values("Content-Type"))

	contentType := fileHeader.Header.Values("Content-Type")[0]
	AllowImageTypes := []string{"application/pdf", "application/json", "application/xml", "application/x-xls", "application/msword", "application/zip", "application/gzip", "text/plain"}
	if !common.IsStringItemExist(AllowImageTypes, contentType) {
		_ = ctx.Error(errors.New("Content-Type Is Not Allowed! "))
		return
	}

	// Upload the file to specific dst.
	dst := common.UploadImagesPath + "/" + fileHeader.Filename
	if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {
		_ = ctx.Error(errors.New("Save Upload File Fail "))
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

func (api *ResourceApi) uploadVideoHandler(ctx *gin.Context) {
	var err error
	var fileHeader *multipart.FileHeader
	// var file multipart.File
	if fileHeader, err = ctx.FormFile("file"); err != nil {
		_ = ctx.Error(err)
		return
	}
	fmt.Printf("FileHeader: %#v \n", fileHeader.Header.Values("Content-Type"))

	contentType := fileHeader.Header.Values("Content-Type")[0]
	AllowImageTypes := []string{"video/mpeg4", "video/avi", "video/x-ms-wmv", "video/mpg"}
	if !common.IsStringItemExist(AllowImageTypes, contentType) {
		_ = ctx.Error(errors.New("Content-Type Is Not Allowed! "))
		return
	}

	// Upload the file to specific dst.
	dst := common.UploadImagesPath + "/" + fileHeader.Filename
	if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {
		_ = ctx.Error(errors.New("Save Upload Video Fail "))
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}

func (api *ResourceApi) uploadAudioHandler(ctx *gin.Context) {
	var err error
	var fileHeader *multipart.FileHeader
	// var file multipart.File
	if fileHeader, err = ctx.FormFile("file"); err != nil {
		_ = ctx.Error(err)
		return
	}
	fmt.Printf("FileHeader: %#v \n", fileHeader.Header.Values("Content-Type"))

	contentType := fileHeader.Header.Values("Content-Type")[0]
	AllowImageTypes := []string{"audio/mid", "audio/mp3", "audio/mpegurl"}
	if !common.IsStringItemExist(AllowImageTypes, contentType) {
		_ = ctx.Error(errors.New("Content-Type Is Not Allowed! "))
		return
	}

	// Upload the file to specific dst.
	dst := common.UploadImagesPath + "/" + fileHeader.Filename
	if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {
		_ = ctx.Error(errors.New("Save Upload Audio Fail "))
		return
	}

	ctx.Set(common.ResponseDataKey, nil)
}
