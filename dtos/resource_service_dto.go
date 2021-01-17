package dtos

type GetQiniuOssClientUploadTokenDTO struct {
	Id uint `validate:"required,numeric" json:"id"`
}

type UploadImageDTO struct {
	Id uint `validate:"required,numeric" json:"id"`
}
type UploadFileDTO struct {
	Id uint `validate:"required,numeric" json:"id"`
}
type UploadVideoDTO struct {
	Id uint `validate:"required,numeric" json:"id"`
}
type UploadAudioDTO struct {
	Id uint `validate:"required,numeric" json:"id"`
}
