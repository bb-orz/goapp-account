package services

import "goapp/dtos"

/* 定义邮件服务模块的服务层方法，并定义数据传输对象DTO*/
var resourceService IResourceService

// 用于对外暴露邮件服务，唯一的暴露点，供接口层调用
func GetResourceService() IResourceService {

	return resourceService
}

// 服务具体实现初始化时设置服务对象，供核心业务层具体实现并设置
func SetResourceService(service IResourceService) {
	resourceService = service
}

type IResourceService interface {
	GetQiniuOssClientUploadToken(dto dtos.GetQiniuOssClientUploadTokenDTO) error // 获取七牛oss客户端上传token
	UploadImage(dto dtos.UploadImageDTO) error                                   // 上传图片到服务器
	UploadFile(dto dtos.UploadFileDTO) error                                     // 上传文件到服务器
	UploadVideo(dto dtos.UploadVideoDTO) error                                   // 上传视频到服务器
	UploadAudio(dto dtos.UploadAudioDTO) error                                   // 上传音频到服务器

}
