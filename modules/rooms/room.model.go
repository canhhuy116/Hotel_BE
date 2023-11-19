package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
)

const RoomEntityName = "Room"

type Room struct {
	bases.BaseModel `json:"-" gorm:"embedded"`
	Name            string         `json:"name" gorm:"column:name; not null; unique"`
	RoomTypeId      int            `json:"room_type_id" gorm:"column:room_type_id; not null"`
	OccupancyStatus string         `json:"occupancy_status" gorm:"column:occupancy_status; default:'available'"`
	Description     string         `json:"description"`
	Images          *common.Images `json:"images" gorm:"column:images; type:json"`
	RoomType        *RoomType      `json:"room_type" gorm:"foreignKey:RoomTypeId"`
}

func (r Room) TableName() string {
	return "rooms"
}

func (r *Room) Mask() {
	r.FakeId(common.DbRoom)
}
