package Models

// type Amn struct {
// 	Business_Centre string `json:"Business_Centre"`
// 	Fitness_Room    string `json:"Fitness_Room"`
// 	Pet_Friendly    string `json:"Pet_Friendly"`
// 	Disabled_Access string `json:"Disabled_Access"`
// 	Air_Conditioned string `json:"Air_Conditioned"`
// 	Free_WIFI       string `json:"Free_WIFI"`
// 	Elevator        string `json:"Elevator"`
// 	Parking         string `json:"Parking"`
// }

type Hotel struct {
	Hotel_Id    string `json:"hotel_id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Address     string `json:"address"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Description string `json:"description"`
	Room_count  int    `json:"room_count"`
	Currency    int    `json:"currency"`
	Amenities   struct {
		Business_Centre string `json:"Business_Centre"`
		Fitness_Room    string `json:"Fitness_Room"`
		Pet_Friendly    string `json:"Pet_Friendly"`
		Disabled_Access string `json:"Disabled_Access"`
		Air_Conditioned string `json:"Air_Conditioned"`
		Free_WIFI       string `json:"Free_WIFI"`
		Elevator        string `json:"Elevator"`
		Parking         string `json:"Parking"`
	}
}

func (b *Hotel) TableName() string {
	return "hotel"
}
