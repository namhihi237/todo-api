package v1

import (
	"fmt"
	"net/http"
	"time"
	"todo/pkg/app"
	"todo/pkg/errors"
	"todo/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" valid:"Required"`
	Email     string     `json:"email" valid:"Required"`
	Password  string     `json:"password" valid:"Required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeleteAt  *time.Time `json:"delete_at"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"Required"`
	Password string `json:"password" valid:"Required"`
}

// @Summary Register user
// @Produce  json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		valid := validation.Validation{}
		var user User

		if err := c.ShouldBind(&user); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		ok, _ := valid.Valid(&user)
		if !ok {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		var userExist User
		if err := db.First(&userExist, "email = ?", user.Email).Error; err == nil {
			appG.Response(http.StatusBadRequest, errors.EMAIL_EXIST, nil)
			return
		}

		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.ERROR_HASH_PASSWORD, nil)
			return
		}
		user.Password = hashedPassword

		if err := db.Create(&user).Error; err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, nil)
	}
}

// @Summary Login
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		valid := validation.Validation{}
		var user UserLogin

		if err := c.ShouldBind(&user); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		// validation email and password
		ok, _ := valid.Valid(&user)
		if !ok {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		var userExist User
		if err := db.First(&userExist, "email = ?", user.Email).Error; err != nil {
			appG.Response(http.StatusUnauthorized, errors.UNAUTHORIZED, nil)
			return
		}

		matchPassword := util.ComparePassword(userExist.Password, user.Password)
		fmt.Println(matchPassword)
		if !matchPassword {
			appG.Response(http.StatusUnauthorized, errors.UNAUTHORIZED, nil)
			return
		}

		fmt.Println("generated token", userExist.Id)
		fmt.Printf("var1 = %T\n", userExist.Id)

		// generated token using id and email
		token, err := util.GenerateToken(userExist.Id, userExist.Email)

		if err != nil {
			fmt.Println(err)
			appG.Response(http.StatusInternalServerError, errors.GENERATED_TOKEN_ERROR, nil)
			return
		}

		fmt.Println("token")
		fmt.Println(token)

		userResponse := UserResponse{
			Id:    userExist.Id,
			Name:  userExist.Name,
			Email: userExist.Email,
		}
		fmt.Println(userResponse)

		appG.Response(http.StatusOK, errors.SUCCESS, map[string]interface{}{
			"token": token,
			"user":  userResponse,
		})

	}
}
