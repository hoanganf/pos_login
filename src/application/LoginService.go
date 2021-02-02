package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/entity/exception"
	"github.com/hoanganf/pos_domain/service"
	"github.com/hoanganf/pos_login/src/application/resource"
	"log"
	"net/http"
)

const (
	PARAM_FROM_URL = "frm"
)

type LoginService struct {
	HomePage    string
	Domains     []string
	TokenName   string
	UserService *service.UserService
}

func NewLoginService(homePage string, domains []string, tokenName string, userService *service.UserService) *LoginService {
	return &LoginService{HomePage: homePage, Domains: domains, TokenName: tokenName, UserService: userService}
}

func (s *LoginService) GetLogin(c *gin.Context) {
	cookie, err := c.Cookie(s.TokenName)
	fromURL := c.DefaultQuery(PARAM_FROM_URL, s.HomePage)
	resource := &resource.Resource{FromURL: fromURL}
	if err != nil {
		//cookie is NotSet
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
		return
	}
	user, cErr := s.UserService.GetUserByCookie(cookie)
	if cErr != nil {
		resource.SetErrorMessage("Cookie is invalid.")
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
		return
	}

	if user.Cookie != "" {
		resource.SetRedirect(fromURL)
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
		return
	}
	resource.SetErrorMessage(exception.GetErrorMessage(exception.CodeSystemError))
	log.Print("jwt not created!")
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})

}

func (s *LoginService) GetLogout(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"resource": &resource.Resource{
			DisableLoader: true,
			FromURL:       c.DefaultQuery(PARAM_FROM_URL, s.HomePage),
			Domains:       s.Domains},
	})
}

func (s *LoginService) Post(c *gin.Context) {
	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	remember := c.PostForm("remember")
	fromURL := c.DefaultPostForm(PARAM_FROM_URL, s.HomePage)
	resource := &resource.Resource{FromURL: fromURL}

	if userName == "" || password == "" {
		//cookie is NotSet
		log.Print(userName, password)
		resource.SetErrorMessage("userName and password is required.")
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
		return
	}
	user, err := s.UserService.GetUserByUserNameAndPassword(userName, password)
	if err != nil {
		resource.SetErrorMessage(err.ErrorMessage)
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
		return
	}

	if user.Cookie != "" {
		if remember != "" {
			resource.SetAccessToken(user.Cookie)
			resource.SetDomains(s.Domains)
		}
		resource.SetRedirect(fromURL)
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
		return
	}
	log.Print("jwt not created!")
	resource.SetErrorMessage(exception.GetErrorMessage(exception.CodeSystemError))
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"resource": resource})
}
