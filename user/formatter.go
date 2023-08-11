package user

type UserFormatter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Token       string `json:"token,omitempty"`
	ImageURL    string `json:"image_url"`
	Role        Role   `json:"role"`
	CreatedAt   string `json:"created_at"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Token:       token,
		ImageURL:    user.AvatarFileName,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return formatter
}

func FormatUsers(users []User) []UserFormatter {
	formatters := []UserFormatter{}
	for _, user := range users {
		formatters = append(formatters, FormatUser(user, ""))
	}

	return formatters
}
