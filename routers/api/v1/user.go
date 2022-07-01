package v1

import (
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

// @Summary Register user
// @Produce  json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
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
