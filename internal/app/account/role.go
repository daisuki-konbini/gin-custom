package account

import "github.com/jinzhu/gorm"

//Role ...
type Role struct {
	Name   string  `gorm:"type:varchar(25);not null;default:'';unique_index:udx_name_domain"`
	Domain string  `gorm:"type:varchar(120);not null;default:'';unique_index:udx_name_domain"`
	Users  []*User `gorm:"many2many:user_roles"`
	gorm.Model
}

//AddPolicy ...
func (s *service) AddPolicy() (err error) {
	// _, err = GetEnforcer().AddPolicySafe(r.Name, r.Domain, r.Object, r.Method)
	return
}
