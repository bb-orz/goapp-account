package dtos

type SendPhoneVerifyCodeDTO struct {
	Phone  uint `validate:"required,numeric" json:"phone" form:"phone"`
	VcType uint `validate:"required,numeric,oneof=1 2 3" json:"email" form:"vctype"`
}
