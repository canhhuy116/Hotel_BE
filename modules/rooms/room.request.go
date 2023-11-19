package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
)

type RoomCreate struct {
	bases.BaseModel `json:",inline" swaggerignore:"true"`
	Name            string         `json:"name" binding:"required"`
	RoomTypeFakeId  string         `json:"room_type_id" binding:"required" gorm:"-"`
	RoomTypeId      int            `json:"-" swaggerignore:"true" `
	Description     string         `json:"description"`
	Images          *common.Images `json:"images"`
}

func (RoomCreate) TableName() string {
	return Room{}.TableName()
}

type RoomUpdate struct {
	Name           string         `json:"name"`
	RoomTypeFakeId string         `json:"room_type_id" gorm:"-"`
	RoomTypeId     int            `json:"-" swaggerignore:"true" `
	Description    string         `json:"description"`
	Images         *common.Images `json:"images"`
}

func (RoomUpdate) TableName() string {
	return Room{}.TableName()
}
