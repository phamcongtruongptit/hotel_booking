package model

import "time"

type BookedRoom struct {
	ID       uint      `gorm:"primary_key;auto_increment" json:"id"`
	CheckIn  time.Time `gorm:"column:checkin;type:date" json:"check_in"`
	CheckOut time.Time `gorm:"column:checkout;type:date" json:"check_out"`
	SellOff  int       `gorm:"column:selloff" json:"sell_off"`

	BookingId uint `gorm:"booking_id"`
	Booking   *Booking

	RoomId uint `gorm:"room_id" json:"room_id"`
	Room   *Room
}

func (*BookedRoom) TableName() string {
	return "booked_room"
}
