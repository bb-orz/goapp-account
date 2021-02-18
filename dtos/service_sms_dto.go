package dtos

type SendPhoneVerifyCodeDTO struct {
	Phone  uint `json:"phone" form:"phone" validate:"required,numeric"`
	VcType uint `json:"vctype" form:"vctype" validate:"required,numeric,oneof=1 2"`
}
