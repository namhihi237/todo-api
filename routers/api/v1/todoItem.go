package v1

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo/pkg/app"
	"todo/pkg/errors"
)

type TodoItem struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Status    string     `json:"status" gorm:"default:'doing';"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeleteAt  *time.Time `json:"delete_at"`
}

// @Summary Get list todo items
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/items [get]
func GetListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		var items []TodoItem
		if err := db.Find(&items).Error; err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, items)
	}
}

// @Summary Create list todo items
// @Produce  json
// @Param title body string true "Title"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/items [post]
func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		var item TodoItem

		//To bind a request body into a type, use model binding.
		//When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header.
		if err := c.ShouldBind(&item); err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		item.Title = strings.TrimSpace(item.Title)
		if item.Title == "" {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		if err := db.Create(&item).Error; err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, item)
	}
}

// @Summary Get a single todo item
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/items/{id} [get]
func GetItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		id, err := strconv.Atoi(c.Param("id")) // convert string to int
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		var item TodoItem
		if err := db.First(&item, id).Error; err != nil {
			appG.Response(http.StatusNotFound, errors.NOT_FOUND, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, item)
	}
}

// @Summary Delete a single todo item
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/items/{id} [delete]
func DeleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		var item TodoItem
		if err := db.First(&item, id).Error; err != nil {
			appG.Response(http.StatusNotFound, errors.NOT_FOUND, nil)
			return
		}

		if err := db.Delete(&item).Error; err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, nil)
	}
}

// @Summary Update a single todo item
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/items/{id} [put]
func UpdateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		var item TodoItem

		if err := c.ShouldBind(&item); err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		if err := db.First(&item, id).Error; err != nil {
			appG.Response(http.StatusNotFound, errors.NOT_FOUND, nil)
			return
		}

		item.Title = strings.TrimSpace(item.Title)
		if item.Title == "" {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAM, nil)
			return
		}

		if err := db.Save(&item).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, item)
	}
}
