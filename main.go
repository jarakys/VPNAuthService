package main

import (
	"VPNAuthService/Configs"
	"VPNAuthService/Controllers"
	"VPNAuthService/DbModels"
	"VPNAuthService/Managers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := Configs.SQLDatabaseConnection("users.db")
	if err != nil {
		panic(err)
	}
	err = db.Table(DbModels.UserModel{}.TableName()).AutoMigrate(&DbModels.UserModel{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	//router.Use(sentrygin.New(sentrygin.Options{}))
	//TODO move PlatformHeaderMiddleware to separate package
	//router.Use(middlewares.PlatformHeaderMiddleware())
	authDbManager := Managers.NewAuthDbManager(db)
	authController := Controllers.NewAuthController(authDbManager)
	router.POST("/register", authController.Register)
	router.POST("/updateUser", authController.Update)
	router.POST("/login", authController.Login)
	router.POST("/deleteUser", authController.Delete)
	router.Run(":8081")
}
