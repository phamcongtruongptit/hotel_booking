package model

import "time"

type Booking struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	BookingDate time.Time `gorm:"booking_date;type:date" json:"booking_date"`
	Note        string    `gorm:"note" json:"note"`
	SellOff     int       `gorm:"selloff" json:"sell_off"`
	ClientName  string    `gorm:"client_name" json:"client_name"`
	UserId      string    `gorm:"user_id" json:"user_id"`
	User        *User

	BookedRooms []*BookedRoom
}

func (*Booking) TableName() string {
	return "booking"
}
