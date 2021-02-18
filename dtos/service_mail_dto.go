package dtos

type SendEmailVerifyCodeDTO struct {
	Email  string `json:"email" form:"email" validate:"required,email"`
	VcType uint   `json:"vctype" form:"vctype" validate:"required,numeric,oneof=1 2 3"`
}
