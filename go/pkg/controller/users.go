package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/domain"
)

type getUser func(uint) (domain.User, error)

type userResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}

type User struct {
	getUser getUser
}

func (u User) GET(c *gin.Context) {
	idUser, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad id",
		})
		return
	}
	user, e := u.getUser(uint(idUser))
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "no exist",
		})
		return
	}

	c.JSON(http.StatusOK, userResponse{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Age:     user.Age,
	})

}

func NewUser(getUser getUser) User {
	return User{
		getUser: getUser,
	}
}
