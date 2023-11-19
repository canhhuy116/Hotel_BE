package rooms

type RoomTypesResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func MapToRoomTypesResponse(roomTypes []RoomType) []RoomTypesResponse {
	roomTypesResponse := make([]RoomTypesResponse, len(roomTypes))

	for i, roomType := range roomTypes {
		roomType.Mask()

		roomTypesResponse[i] = RoomTypesResponse{
			Id:   roomType.FakeID,
			Name: roomType.Name,
		}
	}

	return roomTypesResponse
}
