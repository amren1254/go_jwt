package app

import (
	"go_jwt/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := auth.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
