package account

import (
	"gin-custom/pkg/ecode"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	UserID      string  `gorm:"type:varchar(36);unique_index"`
	Email       string  `gorm:"type:varchar(100);unique;not null"`
	Password    string  `gorm:"type:varchar(100);not null"`
	DisplayName string  `gorm:"type:varchar(30);not null"`
	PhoneNumber string  `gorm:"type:varchar(30);not null;default:''"`
	PhotoURL    string  `gorm:"type:varchar(200);not null;default:''"`
	Roles       []*Role `gorm:"many2many:user_roles"`
	gorm.Model
}

//UserClaims ...
type UserClaims struct {
	UserID string
	Roles  []*Role
	jwt.StandardClaims
}

//Create ...
func (s *service) CreateUser(user *User, roleName string) (*User, error) {
	var role Role
	if err := s.orm.Where("name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}
	user.Roles = append(user.Roles, &role)
	if err := s.orm.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//GetFrom ...
func (s *service) GetFrom(user *User) (*User, error) {
	var ret User
	if err := s.orm.Where(user).First(&ret).Error; err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *service) generageToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.config.GetString("jwt.key")))
	return ss, err
}

//GetToken ...
func (s *service) GetToken(u *User) (string, error) {
	//TODO Regardless of safety
	claims := UserClaims{
		u.UserID,
		u.Roles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(s.config.GetInt64("jwt.expire"))).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	return s.generageToken(claims)
}

//ParseToken ...
func (s *service) ParseToken(tokenStr string) (claims *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.config.GetString("jwt.key"), nil
	})
	if token.Valid {
		if claims, ok := token.Claims.(*UserClaims); ok {
			return claims, nil
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, ecode.TokenExpired
		} else {
			return nil, ecode.TokenErr
		}
	} else {
		return nil, ecode.TokenErr
	}
	return nil, ecode.TokenErr
}
