package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartService struct{}
type CartItem struct {
	ProductID  uint    `json:"product_id"`
	Quantity   float32 `json:"quantity"`
	ItemPrice  float32 `json:"item_price"`
	TotalPrice float32 `json:"total_price"`
}

type UserCart struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`
}

type Cart struct {
	UserCart []UserCart
}

var cart *Cart

func (cs *CartService) AddItem(c *gin.Context) {
	var cartItem CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// var product Product
	// resp, _ := http.Get("https://fakestoreapi.com/product/" + fmt.Sprint(cartItem.ProductID))
	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// json.Unmarshal(body, &product)
	cartItem.TotalPrice = cartItem.ItemPrice * cartItem.Quantity
	cart.UserCart[0].CartItems = append(cart.UserCart[0].CartItems, cartItem)
	c.JSON(http.StatusBadRequest, gin.H{"cart": cart})
}

func (cs *CartService) RemoveItem(c *gin.Context) {
	var cartItem CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var refreshCart []CartItem
	for i, item := range cart.UserCart[0].CartItems {
		if item.ProductID == cartItem.ProductID {
			refreshCart = cart.UserCart[0].CartItems[:i]
			refreshCart = append(refreshCart, cart.UserCart[0].CartItems[i+1:]...)
			break
		}
	}
	cart.UserCart[0].CartItems = refreshCart
	c.JSON(http.StatusBadRequest, gin.H{"cart": cart})
}

func (cs *CartService) Get(c *gin.Context) {
	userIDStr := c.Param("userid")
	for _, userCart := range cart.UserCart {
		if userCart.UserID == c.GetUint(userIDStr) {
			c.JSON(http.StatusOK, userCart)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user cart not found"})
}
