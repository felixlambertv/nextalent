package controller

import (
	"github.com/felixlambertv/nextalent/model"
	"github.com/felixlambertv/nextalent/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PersonController struct {
	DB *gorm.DB
}

func (p *PersonController) GetCountryByPersonName(c *gin.Context) {
	name := c.Param("name")
	var person model.Person
	if result := p.DB.Where("name = ?", name).First(&person); result.Error != nil {
		c.JSON(http.StatusOK, response.BaseResponse{
			Message: "Person not found",
			Status:  false,
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Success get country name",
		Status:  true,
		Data:    person.Country,
	})
}
