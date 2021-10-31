package app

import (
	"go_jwt/auth"
	ent "go_jwt/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

//A sample user
var user = ent.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func Login(c *gin.Context) {
	var u ent.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := auth.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
