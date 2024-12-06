package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/samber/do"
	"gorm.io/gorm"
	"hotel_booking/dto"
	"hotel_booking/model"
	"hotel_booking/repository"
	"net/http"
	"time"
)

type RooomService interface {
	Create(ctx context.Context, room *dto.CreateRoomReq) (*dto.RoomResponse, error)
	Update(ctx context.Context, room dto.Room) (*dto.RoomResponse, error)
	GetRoomById(ctx context.Context, roomId uint) (*dto.RoomResponse, error)
	Delete(ctx context.Context, roomId int) (bool, error)
	List(ctx context.Context, keyword string) (*dto.ListRoomResponse, error)
	SearchFreeRoom(ctx context.Context, checkIn time.Time, checkOut time.Time) (*dto.ListRoomResponse, error)
	GetBookings(ctx context.Context, startTime time.Time, endTime time.Time) (*dto.ListBookingsResponse, error)
	CreateBooking(ctx context.Context, req *model.Booking) (*dto.CreateBookingResponse, error)
}

type roomServiceImpl struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(di *do.Injector) (RooomService, error) {
	return &roomServiceImpl{
		roomRepo: do.MustInvoke[repository.RoomRepository](di),
	}, nil
}

func (s roomServiceImpl) Create(ctx context.Context, req *dto.CreateRoomReq) (*dto.RoomResponse, error) {
	existedRoom, _ := s.roomRepo.FindByName(ctx, req.Name)
	if existedRoom != nil && existedRoom.Name != "" {
		return nil, errors.New(fmt.Sprintf("room with name %s already existed", req.Name))
	}

	room := &model.Room{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Type:        req.Type,
	}
	if err := s.roomRepo.Create(ctx, room); err != nil {
		return nil, err
	}
	roomResponse := &dto.RoomResponse{
		Meta: &dto.Meta{Code: http.StatusOK, Message: "success"},
		Data: room,
	}
	return roomResponse, nil
}

func (s roomServiceImpl) Update(ctx context.Context, room dto.Room) (*dto.RoomResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s roomServiceImpl) GetRoomById(ctx context.Context, roomId uint) (*dto.RoomResponse, error) {
	room, err := s.roomRepo.FindByID(ctx, roomId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with id: %d", roomId)
		}
		return nil, err
	}
	res := &dto.RoomResponse{
		Meta: &dto.Meta{Code: http.StatusOK, Message: "success"},
		Data: room,
	}
	return res, nil
}

func (s roomServiceImpl) Delete(ctx context.Context, roomId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s roomServiceImpl) List(ctx context.Context, keyword string) (*dto.ListRoomResponse, error) {
	rooms, err := s.roomRepo.List(ctx, keyword)
	if err != nil {
		return nil, err
	}
	res := &dto.ListRoomResponse{
		Meta: &dto.Meta{Code: http.StatusOK, Message: "success"},
		Data: rooms,
	}
	return res, err
}

func (s roomServiceImpl) SearchFreeRoom(ctx context.Context, checkIn time.Time, checkOut time.Time) (*dto.ListRoomResponse, error) {
	rooms, err := s.roomRepo.SearchFreeRoom(ctx, checkIn, checkOut)
	if err != nil {
		return nil, err
	}
	res := &dto.ListRoomResponse{
		Meta: &dto.Meta{Code: http.StatusOK, Message: "success"},
		Data: rooms,
	}
	return res, err
}

func (s roomServiceImpl) GetBookings(ctx context.Context, startTime time.Time, endTime time.Time) (*dto.ListBookingsResponse, error) {
	bookings, err := s.roomRepo.GetBooking(ctx, startTime, endTime)
	if err != nil {
		return nil, err
	}
	res := &dto.ListBookingsResponse{
		Meta: &dto.Meta{Code: http.StatusOK, Message: "success"},
		Data: bookings}
	return res, err
}

func (s roomServiceImpl) CreateBooking(ctx context.Context, req *model.Booking) (*dto.CreateBookingResponse, error) {
	err := s.roomRepo.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	res := &dto.CreateBookingResponse{
		Meta: &dto.Meta{Code: http.StatusOK, Message: "success"},
	}
	return res, err

}
