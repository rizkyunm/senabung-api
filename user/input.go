package user

type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"email,required"`
	Password    string `json:"password" binding:"required"`
	Role        Role   `json:"role" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     Role
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
