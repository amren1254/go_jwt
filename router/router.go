package router

import (
	"go_jwt/app"
	"go_jwt/auth"
	"go_jwt/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func InitRouter() {
	router.POST("/login", app.Login)
	router.POST("/todo", middleware.TokenAuthMiddleware(), app.CreateTodo)
	router.POST("/logout", middleware.TokenAuthMiddleware(), app.Logout)
	router.POST("/token/refresh", auth.RefreshToken)

	log.Fatal(router.Run(":8080"))
}
