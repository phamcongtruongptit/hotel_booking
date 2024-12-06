package model

type Room struct {
	//gorm.Model
	ID          uint          `gorm:"primary"`
	Name        string        `gorm:"column:name"`
	Description string        `gorm:"column:description"`
	Price       int           `gorm:"column:price"`
	Type        string        `gorm:"column:type"`
	BookedRoom  []*BookedRoom `gorm:"foreigner:RoomId"`
}

func (*Room) TableName() string {
	return "room"
}
