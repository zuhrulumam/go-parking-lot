package handler

import "github.com/zuhrulumam/doit-test/business/entity"

type ParkResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	SpotID  string `json:"spot_id,omitempty"`
}

type UnparkResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type AvailableSpotResponse struct {
	Success        bool               `json:"success"`
	Message        string             `json:"message,omitempty"`
	VehicleType    string             `json:"vehicle_type"`
	AvailableSpots []ParkingSpotBrief `json:"available_spots"`
}

type ParkingSpotBrief struct {
	SpotID string `json:"spot_id"`
	Floor  int    `json:"floor"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
}

type SearchVehicleResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Vehicle *entity.Vehicle `json:"vehicle,omitempty"`
}

type ErrorResponse struct {
	Success    bool   `json:"success"`
	HumanError string `json:"human_error"`
	DebugError string `json:"debug_error"`
}
