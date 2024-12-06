package model

type Role struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	RoleName string `gorm:"unique;not null"`
	RoleUser []*RoleUser
}
