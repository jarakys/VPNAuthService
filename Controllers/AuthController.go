package Controllers

import (
	"VPNAuthService/APIModels"
	"VPNAuthService/DbModels"
	"VPNAuthService/Managers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewAuthController(db Managers.AuthDbManager) AuthController {
	return &authControllerImpl{db: db}
}

type authControllerImpl struct {
	db Managers.AuthDbManager
}

func (ac *authControllerImpl) Register(c *gin.Context) {
	var user *DbModels.UserModel
	var err error
	if user, err = ac.db.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ac *authControllerImpl) Login(c *gin.Context) {
	id := c.PostForm("id")
	if _, err := ac.db.Get(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ac *authControllerImpl) Update(c *gin.Context) {
	var user APIModels.UserAPIRequestModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ac.db.Update(DbModels.UserModel{
		Id:        user.Id,
		IsPremium: user.IsPremium,
		LastVisit: time.Now().Unix(),
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ac *authControllerImpl) Delete(c *gin.Context) {
	id := c.PostForm("id")
	if err := ac.db.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
