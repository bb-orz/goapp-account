package dtos

type SendPhoneVerifyCodeDTO struct {
	ID    uint `validate:"required,numeric";json:"id"`
	Phone uint `validate:"required,numeric";json:"phone"`
}


