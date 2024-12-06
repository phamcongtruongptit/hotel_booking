package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"hotel_booking/dto"
	"hotel_booking/model"
	"hotel_booking/service"
	"net/http"
	"strconv"
	"time"
)

type RoomController interface {
	GetRoomById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	SearchFreeRoom(ctx *gin.Context)
	GetBookings(ctx *gin.Context)
	CreateBooking(ctx *gin.Context)
}

type roomContr struct {
	roomService service.RooomService
}

func NewRoomController(di *do.Injector) RoomController {
	return &roomContr{
		roomService: do.MustInvoke[service.RooomService](di),
	}
}

func (c roomContr) GetRoomById(ctx *gin.Context) {
	roomId := ctx.Param("room_id")
	uid, _ := strconv.ParseInt(roomId, 10, 64)
	resp, err := c.roomService.GetRoomById(ctx, uint(uid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c roomContr) Create(ctx *gin.Context) {
	req := &dto.CreateRoomReq{}
	_ = ctx.ShouldBind(req)
	resp, err := c.roomService.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c roomContr) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c roomContr) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	resp, err := c.roomService.List(ctx, keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c roomContr) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c roomContr) SearchFreeRoom(ctx *gin.Context) {
	checkInStr := ctx.Query("check_in")
	checkOutStr := ctx.Query("check_out")

	checkIn, err := time.Parse("2006-01-02", checkInStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid check_in date format"})
		return
	}

	checkOut, err := time.Parse("2006-01-02", checkOutStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid check_out date format"})
		return
	}

	resp, err := c.roomService.SearchFreeRoom(ctx, checkIn, checkOut)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c roomContr) GetBookings(ctx *gin.Context) {
	startTimeStr := ctx.Query("start_time")
	endTimeStr := ctx.Query("end_time")
	startTime, err := time.Parse("2006-01-02", startTimeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid check_in date format"})
		return
	}
	endTime, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid check_out date format"})
		return
	}

	resp, err := c.roomService.GetBookings(ctx, startTime, endTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c roomContr) CreateBooking(ctx *gin.Context) {
	req := &model.Booking{}
	_ = ctx.ShouldBind(req)
	resp, err := c.roomService.CreateBooking(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, resp)
}
