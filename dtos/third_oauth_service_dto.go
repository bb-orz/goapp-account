package dtos

type QQLoginDTO struct {
	AccessCode string `validate:"required" json:"access_code"`
}

type WechatLoginDTO struct {
	AccessCode string `validate:"required" json:"access_code"`
}

type WeiboLoginDTO struct {
	AccessCode string `validate:"required" json:"access_code"`
}

type QQBindingDTO struct {
	Id         uint   `validate:"required,numeric" json:"id"`
	AccessCode string `validate:"required" json:"access_code"`
}

type WechatBindingDTO struct {
	Id         uint   `validate:"required,numeric" json:"id"`
	AccessCode string `validate:"required" json:"access_code"`
}

type WeiboBindingDTO struct {
	Id         uint   `validate:"required,numeric" json:"id"`
	AccessCode string `validate:"required" json:"access_code"`
}
