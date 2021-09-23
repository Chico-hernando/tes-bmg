package controllers

import (
	"bmg/configs"
	"bmg/lib/database"
	"bmg/middlewares"
	"bmg/models/user"
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ReffCode(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func CreateUserController(c echo.Context) error {
	var userCreate user.UserCreate
	c.Bind(&userCreate)
	hashedPassword, err := HashPassword(userCreate.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Register User",
			err.Error(),
		))
	}

	var userDB user.User
	userDB.Username = userCreate.Username
	userDB.Email = userCreate.Email
	userDB.Name = userCreate.Name
	userDB.Password = string(hashedPassword)
	userDB.RefferalCode = ReffCode(4)

	e := configs.DB.Create(&userDB).Error
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Register User",
			e.Error(),
		))
	}

	var userResponse = user.UserResponse{
		Id:           userDB.Id,
		Username:     userDB.Username,
		Name:         userDB.Name,
		Email:        userDB.Email,
		RefferalCode: userDB.RefferalCode,
		CreatedAt:    userDB.CreatedAt,
		UpdatedAt:    userDB.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Register User",
		userResponse,
	))
}

func LoginUserController(c echo.Context) error {
	var userLogin user.UserLogin
	c.Bind(&userLogin)

	var userDB user.User
	configs.DB.First(&userDB, "username", userLogin.Username)
	hashedPassword := userDB.Password

	err := VerifyPassword(hashedPassword, userLogin.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Login",
			"Wrong Username or Password",
		))
	}

	token, _ := middlewares.GenerateJWT(userDB.Id)

	var userResponse = user.LoginResponse{
		Username: userDB.Username,
		Name:     userDB.Name,
		Email:    userDB.Email,
		Token:    token,
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Login",
		userResponse,
	))
}

func UpdateUserController(c echo.Context) error {
	var userUpdate user.UserUpdate
	userId, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&userUpdate)

	hashedPassword, _ := HashPassword(userUpdate.Password)

	var userDB user.User
	configs.DB.First(&userDB, "id", userId)
	userDB.Name = userUpdate.Name
	userDB.Email = userUpdate.Email
	userDB.Username = userUpdate.Username
	userDB.Password = string(hashedPassword)

	err := configs.DB.Save(&userDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Data",
			err.Error(),
		))
	}

	var userResponse = user.UserResponse{
		Id:                userDB.Id,
		Username:          userDB.Username,
		Name:              userDB.Name,
		Email:             userDB.Email,
		RefferalCode:      userDB.RefferalCode,
		RefferalCodeInput: userDB.RefferalCodeInput,
		CreatedAt:         userDB.CreatedAt,
		UpdatedAt:         userDB.UpdatedAt,
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data",
		userResponse,
	))
}

func ReffCodeController(c echo.Context) error {
	userId := middlewares.GetUserIdFromJWT(c)

	var userDB user.User
	var userModel user.User
	var reffCode user.UserReff

	c.Bind(&reffCode)

	userReff, err := database.CheckUserByReffCode(reffCode.RefferalCodeInput, userDB)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"Failed get Refferal Code",
			err.Error(),
		))
	}

	if userReff {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Already submit refferal code",
			"",
		))
	}

	error := configs.DB.First(&userDB, "refferal_code", reffCode.RefferalCodeInput).Error

	if error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"Refferal Code not found",
			"",
		))
	}

	if userDB.Id == userId {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Can't use your own refferal code",
			"",
		))
	}

	var userResponse user.UserResponse
	configs.DB.Model(userModel).First(&userResponse, "id", userId)
	userResponse.RefferalCodeInput = reffCode.RefferalCodeInput
	er := configs.DB.Model(userModel).Where("id = ?", userId).Update("refferal_code_input", reffCode.RefferalCodeInput).Error

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Input Refferal Code",
			er.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Input Refferal Code",
		userResponse,
	))
}

func GetUserByName(c echo.Context) error {
	var userData []user.UserResponse
	var err error

	name := c.QueryParam("name")

	userData, err = database.GetUserByName(name)

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		userData,
	))
}
