package forms

//SigninForm ...
type SigninForm struct {
	JsonForm
	ID       uint   `form:"id" json:"id"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

//SignupForm ...
type SignupForm struct {
	JsonForm
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
