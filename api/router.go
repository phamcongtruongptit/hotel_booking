package api

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"hotel_booking/controller"
)

func InitRouter(di *do.Injector) (*gin.Engine, error) {
	//cf := do.MustInvoke[*conf.Config](di)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	roomController := controller.NewRoomController(di)

	v1 := r.Group("/api/v1")

	roomGroup := v1.Group("/room")
	roomGroup.GET("", roomController.List)
	roomGroup.GET("/:room_id", roomController.GetRoomById)
	roomGroup.POST("", roomController.Create)
	roomGroup.PUT("/:id", roomController.Update)
	roomGroup.GET("/free-room", roomController.SearchFreeRoom)
	roomGroup.DELETE("/user/:id", roomController.Delete)

	bookingGroup := v1.Group("/booking")
	bookingGroup.POST("/create", roomController.CreateBooking)
	bookingGroup.GET("", roomController.GetBookings)
	return r, nil
}
