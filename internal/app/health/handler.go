package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Service ...
type Service interface {
	Check(c *gin.Context)
}

type service struct {
	orm *gorm.DB
}

//NewService ...
func NewService(orm *gorm.DB) Service {
	return &service{
		orm: orm,
	}
}

func (s *service) Check(c *gin.Context) {
	err := s.orm.DB().Ping()
	if err != nil {
		c.String(http.StatusBadRequest, "error")
	}
	c.String(http.StatusOK, "ok")
}
