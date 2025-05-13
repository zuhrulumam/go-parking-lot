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
	UseLock     bool        `json:"use_lock"`
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
	ID            uint       `gorm:"primaryKey" json:"id"`
	VehicleNumber string     `json:"vehicle_number"`
	VehicleType   string     `gorm:"size:1" json:"vehicle_type"` // 'B', 'M', 'A'
	SpotID        string     `json:"spot_id"`
	ParkedAt      time.Time  `json:"parked_at"`
	UnparkedAt    *time.Time `json:"unparked_at"`
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
}

type UpdateParkingSpot struct {
	ID       uint `json:"id"`
	Floor    int
	Row      int
	Col      int
	Occupied *bool `json:"occupied"`
}

type InsertVehicle struct {
	VehicleNumber string
	VehicleType   string
	SpotID        string
}

type UpdateVehicle struct {
	ID            uint
	VehicleNumber string
	UnparkedAt    *time.Time
}

type SpotID struct {
	Floor int
	Row   int
	Col   int
}
