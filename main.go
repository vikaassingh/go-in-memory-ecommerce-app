package main

import (
	"go-in-memory-ecommerce-app/api"

	"github.com/gin-gonic/gin"
)

var userService api.UserService
var productService api.ProductService
var cartService api.CartService
var orderService api.OrderService

func main() {
	r := gin.Default()
	// r.Use()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// UserService
	r.GET("/api/user/:id", userService.Get)
	r.GET("/api/user", userService.GetAll)

	// ProductService
	r.GET("/api/product/:id", productService.Get)
	r.GET("/api/product", productService.GetAll)

	// CartService
	r.PATCH("/api/cart/add", cartService.AddItem)
	r.PATCH("/api/cart/remove", cartService.RemoveItem)
	r.GET("/api/cart/:userid", cartService.Get)

	// OrderService
	r.GET("/api/order/:id", orderService.Get)
	r.GET("/api/order", orderService.GetAll)
	r.POST("/api/order/create", orderService.Create)

	r.Run(":5000")
}
