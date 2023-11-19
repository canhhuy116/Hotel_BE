package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
)

const RoomTypeEntityName = "RoomType"

type RoomType struct {
	bases.BaseModel        `json:"-" gorm:"embedded"`
	Name                   string  `json:"name" gorm:"column:name; not null; unique"`
	BedCount               int     `json:"bed_count" gorm:"column:bed_count"`
	ChargesForCancellation float64 `json:"charges_for_cancellation" gorm:"column:charges_for_cancellation"`
	Price                  float64 `json:"price" gorm:"column:price"`
	FoodOption             string  `json:"food_option" gorm:"column:food_option"`
}

func (r RoomType) TableName() string {
	return "room_types"
}

func (r *RoomType) Mask() {
	r.FakeId(common.DbRoomType)
}
