package dtos

type QQLoginDTO struct {
	AccessCode string `json:"access_code" form:"access_code" validate:"required"`
}

type WechatLoginDTO struct {
	AccessCode string `json:"access_code" form:"access_code" validate:"required"`
}

type WeiboLoginDTO struct {
	AccessCode string `json:"access_code" form:"access_code" validate:"required"`
}

type QQBindingDTO struct {
	Id         uint   `json:"id" validate:"required,numeric"`
	AccessCode string `json:"access_code" validate:"required"`
}

type WechatBindingDTO struct {
	Id         uint   `json:"id" validate:"required,numeric"`
	AccessCode string `json:"access_code" validate:"required"`
}

type WeiboBindingDTO struct {
	Id         uint   `json:"id" validate:"required,numeric"`
	AccessCode string `json:"access_code" validate:"required"`
}
