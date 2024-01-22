package app

import (
	"github.com/AdiLambe/TestGoLangCodes/workspace/domain"
	"github.com/AdiLambe/TestGoLangCodes/workspace/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// create an instance of the repository
	orderRepo := domain.NeworderRepositoryDb()

	// create an instance of the service and inject the repository
	orderService := service.NewOrderService(orderRepo)

	// create an instance of the handlers and inject the services
	orderHandlers := &OrderHandlers{service: orderService}

	//define route
	router.POST("/postOrder", orderHandlers.CreateOrder)      // To Post Order request
	router.GET("/getOrder/:order_id", orderHandlers.GetOrder) // To Get Order Request for particular ID
	router.GET("/orderslist", orderHandlers.GetOrdersList)    // To List all Order

	return router

}

// starting server
func Start() {
	r := SetupRouter()
	r.Run("localhost:8000")
}
