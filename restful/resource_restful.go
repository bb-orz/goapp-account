package restful

import (
	"github.com/bb-orz/goinfras/XGin"
	"github.com/gin-gonic/gin"
	"goapp/restful/middleware"
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

	engine.Static("/image", "/upload/images")
	engine.Static("/file", "/upload/files")
	engine.Static("/video", "/upload/videos")
	engine.Static("/audio", "/upload/audios")

	// 用户鉴权访问路由组接口
	userGroup := engine.Group("/resource", middleware.JwtAuthMiddleware())
	userGroup.GET("/get_qiniu_upload_token", api.getQiniuUploadTokenHandler)
	userGroup.POST("/upload_image", api.uploadImageHandler)
	userGroup.POST("/upload_file", api.uploadFileHandler)
	userGroup.POST("/upload_video", api.uploadVideoHandler)
	userGroup.POST("/upload_audio", api.uploadAudioHandler)

}

func (api *ResourceApi) getQiniuUploadTokenHandler(ctx *gin.Context) {

}

func (api *ResourceApi) uploadImageHandler(ctx *gin.Context) {

}

func (api *ResourceApi) uploadFileHandler(ctx *gin.Context) {

}

func (api *ResourceApi) uploadVideoHandler(ctx *gin.Context) {

}

func (api *ResourceApi) uploadAudioHandler(ctx *gin.Context) {

}
