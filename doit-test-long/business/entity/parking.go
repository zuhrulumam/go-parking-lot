package entity

import "time"

type VehicleType string

const (
	Bicycle    VehicleType = "B"
	Motorcycle VehicleType = "M"
	Automobile VehicleType = "A"
)

type Spot struct {
	SpotType  VehicleType
	IsActive  bool
	Occupied  bool
	VehicleNo string
}

type GetAvailableParkingSpot struct {
	VehicleType VehicleType `json:"vehicle_type"`
	Active      *bool       `json:"active"`
	Occupied    *bool       `json:"occupied"`
}

type ParkingSpot struct {
	ID       uint `gorm:"primaryKey"`
	Floor    int
	Row      int
	Col      int
	Type     string `gorm:"size:1"` // 'B', 'M', 'A', 'X'
	Active   bool
	Occupied bool
}

type Vehicle struct {
	ID            uint `gorm:"primaryKey"`
	VehicleNumber string
	VehicleType   string `gorm:"size:1"` // 'B', 'M', 'A'
	SpotID        string
	ParkedAt      time.Time
	UnparkedAt    *time.Time
}

type Park struct {
	VehicleType   VehicleType `json:"vehicle_type"`
	VehicleNumber string      `json:"vehicle_number"`
}

type UnPark struct {
	SpotID        string `json:"spot_id"`
	VehicleNumber string `json:"vehicle_number"`
}

type GetAvailablePark struct {
	VehicleType VehicleType `json:"vehicle_type"`
}

type SearchVehicle struct {
	VehicleNumber string `json:"vehicle_number"`
	SpotID        string `json:"spot_id"`
}

type UpdateParkingSpot struct {
	Occupied bool `json:"occupied"`
}

type InsertVehicle struct {
	VehicleNumber string
	VehicleType   string
	SpotID        string
}

type UpdateVehicle struct {
	VehicleNumber string
	UnparkedAt    *time.Time
}
