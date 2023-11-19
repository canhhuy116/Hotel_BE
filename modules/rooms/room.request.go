package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
)

type RoomCreate struct {
	bases.BaseModel `json:",inline" swaggerignore:"true"`
	Name            string         `json:"name" binding:"required"`
	RoomTypeId      int            `json:"room_type_id" binding:"required"`
	Description     string         `json:"description"`
	Images          *common.Images `json:"images"`
}

func (RoomCreate) TableName() string {
	return Room{}.TableName()
}
