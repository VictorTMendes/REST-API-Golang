package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	//Usecase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			Id:    1,
			Name:  "Batata Frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
