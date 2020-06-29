package forms

//StudentRegister ...
type StudentRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gt=6"`
}
