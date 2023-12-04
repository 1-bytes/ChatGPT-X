package requests

// ValidateDoRegister 用户注册验证。
type ValidateDoRegister struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=30"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=30"`
}

// ValidateDoLogin 用户登录验证。
type ValidateDoLogin struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=30"`
	Password string `form:"password" binding:"required,min=6,max=30"`
}

// ValidateUserDelete 用户删除验证。
type ValidateUserDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
