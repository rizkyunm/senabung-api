package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rizkyunm/senabung-api/auth"
	"github.com/rizkyunm/senabung-api/helper"
	"github.com/rizkyunm/senabung-api/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// catch input from user
	// mapping input from user to struct RegisterUserInput

	var input user.RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	registerUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(registerUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(registerUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	// user entry input (email & password)
	// input catch by handler
	// mapping data from user input --> input struct
	// passing service
	// Find specified email in service
	// match password

	var input user.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	role := c.Query("role")
	if role != "admin" {
		role = "client"
	}

	input.Role = user.Role(role)

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if loggedInUser.ID == 0 || loggedInUser.Role != input.Role {
		errorMessage := gin.H{"errors": "user not found"}

		response := helper.APIResponse("Login failed", http.StatusNotFound, "error", errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// catch input from user
	// mapping input to struct input
	// pass struct input to service
	// service will call repository - service
	// repository - db
	var input user.CheckEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "Email has been used"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// catch input from user
	// save to image folder ("images")
	// call repo from service
	// use JWT to catch who is upload the avatar
	// catch user ID
	// update user data, save file location

	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println("parse file", err.Error())
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		fmt.Println("save to server", err.Error())
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		fmt.Println("save to db", err.Error())
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("current_user").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		response := helper.APIResponse("Can't get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUsers(users)

	response := helper.APIResponse("List of users", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUser(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		response := helper.APIResponse("Can't get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUsers(users)

	response := helper.APIResponse("List of users", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
