package dto

import (
	"hotel_booking/model"
	"time"
)

type Room struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Type        string `json:"type"`
}

type CreateRoomReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Type        string `json:"type"`
}

type ListRoomResponse struct {
	Meta *Meta         `json:"meta"`
	Data []*model.Room `json:"data"`
}

type ListBookingsResponse struct {
	Meta *Meta            `json:"meta"`
	Data []*model.Booking `json:"data"`
}

type RoomResponse struct {
	Meta *Meta       `json:"meta"`
	Data *model.Room `json:"data"`
}

type CreateBookingResponse struct {
	Meta *Meta `json:"meta"`
}

type BookedRoom struct {
	CheckIn  time.Time `json:"check_in;type:date"`
	CheckOut time.Time `json:"check_out;type:date"`
	SellOff  int       `json:"sell_off;type:int"`
	//BookingId int       `json:"booking_id;type:int"`
	RoomId int `json:"room_id;type:int"`
}

type Booking struct {
	BookingDate time.Time     `json:"booking_date;type:date"`
	Note        string        `json:"note"`
	SellOff     int           `json:"sell_off;type:int64"`
	ClientName  string        `json:"client_name"`
	UserId      int           `json:"user_id;type:int64"`
	BookedRooms []*BookedRoom `json:"booked_rooms"`
}
