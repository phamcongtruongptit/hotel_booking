package repository

import (
	"context"
	"github.com/samber/do"
	"gorm.io/gorm"
	"hotel_booking/model"
	"time"
)

type RoomRepository interface {
	FindByID(ctx context.Context, id uint) (*model.Room, error)
	List(ctx context.Context, keyword string) ([]*model.Room, error)
	Create(ctx context.Context, u *model.Room) error
	Update(ctx context.Context, u *model.Room) error
	FindByName(ctx context.Context, name string) (*model.Room, error)
	SearchFreeRoom(ctx context.Context, checkIn time.Time, checkOut time.Time) ([]*model.Room, error)
	GetBooking(ctx context.Context, startTime time.Time, endTime time.Time) ([]*model.Booking, error)
	CreateBooking(ctx context.Context, req *model.Booking) error
}

func newRoomRepository(di *do.Injector) (RoomRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &roomRepo{db: db}, nil
}

type roomRepo struct {
	db *gorm.DB
}

func (r roomRepo) FindByID(ctx context.Context, id uint) (*model.Room, error) {
	var room model.Room
	err := r.db.First(&room, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &room, err
}

func (r roomRepo) List(ctx context.Context, keyword string) ([]*model.Room, error) {
	var rooms []*model.Room
	err := r.db.WithContext(ctx).
		Model(&model.Room{}).
		Where("name like ? or description like ?", "%"+keyword+"%", "%"+keyword+"%").
		Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r roomRepo) Create(ctx context.Context, room *model.Room) error {
	return r.db.WithContext(ctx).Create(room).Error
}

func (r roomRepo) Update(ctx context.Context, u *model.Room) error {
	//TODO implement me
	panic("implement me")
}

func (r roomRepo) FindByName(ctx context.Context, name string) (*model.Room, error) {
	var room *model.Room
	err := r.db.WithContext(ctx).First(&room, "name = ?", name).Error
	return room, err
}

func (r roomRepo) SearchFreeRoom(ctx context.Context, start time.Time, end time.Time) ([]*model.Room, error) {
	var rooms []*model.Room
	err := r.db.WithContext(ctx).
		Table("room").
		Select("room.*").
		Not("id IN (?)",
			r.db.WithContext(ctx).
				Table("booked_room").
				Select("room_id").
				Where("checkout > ? AND checkin < ?", start, end)).
		Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r roomRepo) GetBooking(ctx context.Context, startTime time.Time, endTime time.Time) ([]*model.Booking, error) {
	var bookings []*model.Booking
	err := r.db.WithContext(ctx).
		Model(&model.Booking{}).
		Preload("User").
		Preload("BookedRooms").
		Where("booking_date >= ? AND booking_date <= ?", startTime, endTime).
		Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r roomRepo) CreateBooking(ctx context.Context, booking *model.Booking) error {
	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, bookedRoom := range booking.BookedRooms {
		if err := tx.Create(&bookedRoom).Table("booked_room").Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
