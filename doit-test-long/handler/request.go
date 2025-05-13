package handler

type ParkRequest struct {
	VehicleType   string `json:"vehicle_type"`
	VehicleNumber string `json:"vehicle_number"`
}

type UnparkRequest struct {
	SpotID        string `json:"spot_id"`
	VehicleNumber string `json:"vehicle_number"`
}
