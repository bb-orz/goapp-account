package dtos

type SendPhoneVerifyCodeDTO struct {
	Id    uint `validate:"required,numeric";json:"id"`
	Phone uint `validate:"required,numeric";json:"phone"`
}
