package account

import (
	"gin-custom/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//Service ...
type Service interface {
	Register(*gin.Context) interface{}
	Login(*gin.Context) interface{}
}

type service struct {
	orm    *gorm.DB
	config *viper.Viper
}

//NewService ...
func NewService(orm *gorm.DB, conf *viper.Viper) Service {
	return &service{
		orm:    orm,
		config: conf,
	}
}

func (s *service) Register(c *gin.Context) interface{} {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,gt=6"`
		DisplayName string `json:"display_name" binding:"required"`
		Role        string `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}
	user, err := s.CreateUser(&User{
		UserID:      uuid.New().String(),
		Email:       req.Email,
		Password:    helpers.Encrypy(req.Password),
		DisplayName: req.DisplayName,
	}, req.Role)
	if err != nil {
		return err
	}
	token, err := s.GetToken(user)
	if err != nil {
		return err
	}
	return token
}

func (s *service) Login(c *gin.Context) interface{} {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gt=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}
	user, err := s.GetFrom(&User{
		Email:    req.Email,
		Password: helpers.Encrypy(req.Password),
	})
	if err != nil {
		return err
	}
	token, err := s.GetToken(user)
	if err != nil {
		return err
	}
	c.Header("token", token)
	return nil
}
