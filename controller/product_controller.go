package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}

}

func (p *productController) GetProduct(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)

}
