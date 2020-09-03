package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/controller"
	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/domain"
	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/service"
	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/usecase"
)

/*
ID
Name
Surname
Age
*/

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	cache := service.NewUserCache()
	cache.PutUser(domain.User{
		ID:      5,
		Name:    "Testing",
		Surname: "Rodriguez",
		Age:     100,
	})

	provider := service.NewUserProvider(cache.GetUser)
	userGETUseCase := usecase.NewUser(provider)
	userController := controller.NewUser(userGETUseCase.GetUser)

	r.GET("/users/:id", userController.GET)
	r.Run()
}
