package forms

//VodForm ...
type VodForm struct {
	JsonForm
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
