package dtos

type SendEmailForVerifyDTO struct {
	Id    uint   `validate:"required,numeric";json:"id"`
	Email string `validate:"required,email";json:"email"`
}

type SendEmailForgetPasswordDTO struct {
	Id    uint   `validate:"required,numeric";json:"id"`
	Email string `validate:"required,email";json:"email"`
}
