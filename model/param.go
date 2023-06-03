package model

type SignupForm struct {
	//Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	//Email      string `json:"email" binding:"required,email"`
}

type LoginForm struct {
	//Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	//Email      string `json:"email" binding:"required,email"`
}
