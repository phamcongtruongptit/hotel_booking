package model

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment"`
	Email    string `gorm:"email, unique"`
	FullName string `gorm:"full_name"`
	Password string `gorm:"password"`
	UserName string `gorm:"user_name, unique"`
	Booking  []*Booking
	RoleUser []*RoleUser
}

func (*User) TableName() string {
	return "user"
}
