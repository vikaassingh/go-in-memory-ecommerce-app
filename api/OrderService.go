package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderService struct{}

// type OrderItem CartItem

type Order struct {
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	OrderNumber string     `json:"order_number"`
	OrderItem   []CartItem `json:"order_item"`
	Amount      float32    `json:"amount"`
}

type Orders []Order

var orders *Orders

func (us *OrderService) Get(c *gin.Context) {
	orderNumber := c.Param("order_number")
	for _, order := range *orders {
		if order.OrderNumber == orderNumber {
			c.JSON(http.StatusOK, order)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
}

func (us *OrderService) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Orders": orders,
	})
}

func (os *OrderService) Create(c *gin.Context) {
	userID := c.GetHeader("user_id")
	orderNumber := generateOrderNumber()
	resp, err := http.Get("127.0.0.1:5000/api/cart/" + userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var userCart UserCart
	json.Unmarshal(body, &userCart)
	order := Order{
		UserID:      userCart.UserID,
		OrderNumber: orderNumber,
		OrderItem:   userCart.CartItems,
	}

	*orders = append(*orders, order)
	c.JSON(http.StatusOK, order)
}

func generateOrderNumber() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number (for example, between 100000 and 999999)
	randomNum := rand.Intn(900000) + 100000

	// Optionally, you can add a prefix or a timestamp
	orderNumber := fmt.Sprintf("ORD-%d-%d", time.Now().Year(), randomNum)

	return orderNumber
}
