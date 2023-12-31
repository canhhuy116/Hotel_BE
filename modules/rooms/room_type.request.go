package rooms

import "Hotel_BE/modules/bases"

type RoomTypeCreate struct {
	bases.BaseModel        `json:",inline" swaggerignore:"true"`
	Name                   string  `json:"name" binding:"required"`
	BedCount               int     `json:"bed_count" binding:"required"`
	ChargesForCancellation float64 `json:"charges_for_cancellation"`
	Price                  float64 `json:"price" binding:"required"`
	FoodOption             string  `json:"food_option" binding:"required"`
}

func (RoomTypeCreate) TableName() string {
	return RoomType{}.TableName()
}

type RoomTypeUpdate struct {
	Name                   string  `json:"name"`
	BedCount               int     `json:"bed_count"`
	ChargesForCancellation float64 `json:"charges_for_cancellation"`
	Price                  float64 `json:"price"`
	FoodOption             string  `json:"food_option"`
}

func (RoomTypeUpdate) TableName() string {
	return RoomType{}.TableName()
}
