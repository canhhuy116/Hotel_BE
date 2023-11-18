package rooms

import "Hotel_BE/modules/bases"

type RoomTypeCreate struct {
	bases.BaseModel        `json:",inline"`
	Name                   string  `json:"name" binding:"required"`
	BedCount               int     `json:"bed_count" binding:"required"`
	ChargesForCancellation float64 `json:"charges_for_cancellation" binding:"required"`
	Price                  float64 `json:"price" binding:"required"`
	FoodOption             string  `json:"food_option" binding:"required"`
}

func (RoomTypeCreate) TableName() string {
	return RoomType{}.TableName()
}
