package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AdiLambe/TestGoLangCodes/workspace/domain"
	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
	"github.com/AdiLambe/TestGoLangCodes/workspace/service"
	"github.com/gin-gonic/gin"
)

type result struct {
	orders []domain.Order
	err    *errs.AppError
}

type OrderHandlers struct {
	service service.OrderService
}

func (ch *OrderHandlers) CreateOrder(c *gin.Context) {
	var newOrder domain.Order

	//Bind the JSON data from the request body to the newOrder struct
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedOrder, err := ch.service.SaveOrder(newOrder)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusCreated, savedOrder)
}

// Handler Func
func (ch *OrderHandlers) GetOrdersList(c *gin.Context) {
	status := c.Query("status")
	orders, err := ch.service.GetOrdersList(status)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (ch *OrderHandlers) GetOrder(c *gin.Context) {

	id := c.Param("order_id")
	log.Printf("Received request for Order ID: %s", id)

	order, err := ch.service.GetOrder(id)
	if err != nil {
		log.Printf("Error: %+v", err)
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func (ch *OrderHandlers) GetOrdersListWithTimeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Minute) // Set a timeout of 1 minutes
	defer cancel()                                                         // Ensure the context is canceled to release resources

	status := c.Query("status")

	// Create a channel to receive the result from a goroutine
	resultChan := make(chan result, 1)

	go func() {
		orders, err := ch.service.GetOrdersList(status)
		resultChan <- result{orders, err}
	}()

	select {
	case <-ctx.Done():
		// If the context is canceled, it means the timeout has occurred
		c.JSON(http.StatusRequestTimeout, gin.H{"message": "Request timed out"})
	case res := <-resultChan:
		// Process the result as usual
		if res.err != nil {
			c.JSON(res.err.Code, gin.H{"message": res.err.Message})
		} else {
			c.JSON(http.StatusOK, res.orders)
		}
	}
}
