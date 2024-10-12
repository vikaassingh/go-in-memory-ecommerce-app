package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductService struct{}

type Product struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func (ps *ProductService) Get(c *gin.Context) {
	var product Product
	id := c.Param("id")
	resp, err := http.Get("https://fakestoreapi.com/product/" + fmt.Sprint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	json.Unmarshal(body, &product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func (ps *ProductService) GetAll(c *gin.Context) {
	var products []Product
	resp, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	json.Unmarshal(body, &products)
	c.JSON(http.StatusOK, gin.H{
		"product": products,
	})
}
