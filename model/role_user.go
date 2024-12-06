package model

type RoleUser struct {
	ID     uint `gorm:"primary_key;AUTO_INCREMENT"`
	RoleId uint `gorm:"role_id"`
	UserId uint `gorm:"user_id"`

	User *User
	Role *Role
}
