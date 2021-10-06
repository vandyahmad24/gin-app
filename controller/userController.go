package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vandyahmad24/gin-app/entity/user"
	"github.com/vandyahmad24/gin-app/helper"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dar user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita passing sebagai parameter ke service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "Belum ada")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// ambil input dari user (email, password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing ke service
	// diservice mencari dengan bantuan repository user dengan email x
	// jika ketemu mencocokan password
	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loginUser, "belum ada token")
	response := helper.APIResponse("Success Login User", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailIsUse(c *gin.Context) {
	// ada input email dari user
	// input email dimapping ke struct input
	// struct input di passwing ke service
	// service akan manggil repository apakah ada atau tidak emailnya
	//  repository db
	var input user.CekEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking errors", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	isEmailAvailable, err := h.userService.CekEmail(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking errors", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user -> berupa form-data
	// simpan gambar difolder "images/"
	// diservice panggil repo
	// repo menentukan siapa user yg mengakses (JWT sementara hardcode)
	// repo ambil data user yg ID nya 1
	// repo update data user simpan lokasi file
	// c.SaveUploadedFile(file)
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 18
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	// path := "images/" + strconv.Itoa(userID) + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// harusnya dpet dari jwt

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar Successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
