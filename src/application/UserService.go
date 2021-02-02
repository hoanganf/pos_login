package application

import (
	"github.com/hoanganf/pos_common/client/resource"
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/entity/exception"
	"github.com/hoanganf/pos_domain/service"
	"net/http"
)

type UserService struct {
	UserService *service.UserService
	UserFactory entity.UserFactory
}

func NewUserService(userService *service.UserService, userFactory entity.UserFactory) *UserService {
	return &UserService{UserService: userService,
		UserFactory: userFactory}
}

func (s *UserService) Login(c *gin.Context) {
	requestBody := resource.LoginRequestResource{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.UserService.GetUser(&entity.LoginInfo{
		UserName: requestBody.UserName,
		Password: requestBody.Password,
		Cookie:      requestBody.Cookie,
	})

	if err != nil {
		if err.ErrorCode == exception.CodeSignatureInvalid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		if err.ErrorCode == exception.CodeNotFound {
			c.JSON(http.StatusNotFound, err)
			return
		}
		if err.ErrorCode == exception.CodeValueInvalid {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, user)
	} else {
		c.JSON(200, s.UserFactory.Create(user, fields))
	}
}
