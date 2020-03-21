package controllers

import (
	"net/http"

	"github.com/devarchi33/all-about-shoes-api/models"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

type ShoesController struct{}

func (c ShoesController) Init(g echoswagger.ApiGroup) {
	g.POST("", c.Create).
		AddParamBody([]models.ShoesCreateDto{}, "[]ShoesCreateDto", "[]ShoesCreateDto", true)
}

func (ShoesController) Create(c echo.Context) error {
	var params []models.ShoesCreateDto
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, param := range params {
		d, err := models.Shoes{}.TranslateForCreate(param)
		if err != err {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := d.Create(c.Request().Context()); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return ReturnApiSucc(c, http.StatusCreated, int64(len(params)), nil)
}
