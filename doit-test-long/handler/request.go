package handler

type ParkRequest struct {
	VehicleType   string `json:"vehicle_type" validate:"required,oneof=M B A"`
	VehicleNumber string `json:"vehicle_number" validate:"required"`
}

type UnparkRequest struct {
	SpotID        string `json:"spot_id"`
	VehicleNumber string `json:"vehicle_number" validate:"required"`
}
