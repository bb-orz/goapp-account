package dtos

type SendEmailForVerifyDTO struct {
	ID    uint   `validate:"required,numeric";json:"id"`
	Email string `validate:"required,email";json:"email"`
}

type SendEmailForgetPasswordDTO struct {
	ID    uint   `validate:"required,numeric";json:"id"`
	Email string `validate:"required,email";json:"email"`
}
