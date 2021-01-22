package dtos

type SendEmailVerifyCodeDTO struct {
	Email  string `validate:"required,email" json:"email" form:"email"`
	VcType uint   `validate:"required,numeric,oneof=1 2 3" json:"email" form:"vctype"`
}
