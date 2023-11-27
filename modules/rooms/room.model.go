package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
)

const RoomEntityName = "Room"

type Room struct {
	bases.BaseModel `json:",inline" gorm:"embedded"`
	Name            string         `json:"name" gorm:"column:name; not null; unique"`
	RoomTypeId      int            `json:"-" gorm:"column:room_type_id; not null"`
	OccupancyStatus string         `json:"occupancy_status" gorm:"column:occupancy_status; default:'available'"`
	Description     string         `json:"description"`
	Images          *common.Images `json:"images" gorm:"column:images; type:json"`
	RoomType        *RoomType      `json:"roomType" gorm:"foreignKey:RoomTypeId; preload:false"`
}

func (r Room) TableName() string {
	return "rooms"
}

func (r *Room) Mask() {
	r.FakeId(common.DbRoom)
	r.RoomType.Mask()
}
