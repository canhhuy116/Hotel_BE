package rooms

type RoomTypeResponse struct {
	Id                     string  `json:"id"`
	Name                   string  `json:"name"`
	Price                  float64 `json:"price"`
	BedCount               int     `json:"bed_count"`
	ChargesForCancellation float64 `json:"charges_for_cancellation"`
	FoodOption             string  `json:"food_option"`
}

func MapToRoomTypeResponse(roomType *RoomType) RoomTypeResponse {
	roomType.Mask()

	roomTypeResponse := RoomTypeResponse{
		Id:                     roomType.FakeID,
		Name:                   roomType.Name,
		Price:                  roomType.Price,
		BedCount:               roomType.BedCount,
		ChargesForCancellation: roomType.ChargesForCancellation,
		FoodOption:             roomType.FoodOption,
	}

	return roomTypeResponse
}
