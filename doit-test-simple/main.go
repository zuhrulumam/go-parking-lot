package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

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

type ParkingLot struct {
	floors      int
	rows        int
	cols        int
	grid        [][][]*Spot
	vehicleMap  map[string]string // vehicleNumber -> spotId
	lastSpotMap map[string]string // last spot history
	mutex       sync.RWMutex
}

func NewParkingLot(floors, rows, cols int, layout [][][]string) *ParkingLot {
	grid := make([][][]*Spot, floors)
	for f := 0; f < floors; f++ {
		grid[f] = make([][]*Spot, rows)
		for r := 0; r < rows; r++ {
			grid[f][r] = make([]*Spot, cols)
			for c := 0; c < cols; c++ {
				cell := layout[f][r][c]
				parts := strings.Split(cell, "-")
				typ := parts[0]
				active := parts[1] == "1"

				var spotType VehicleType
				switch typ {
				case "B":
					spotType = Bicycle
				case "M":
					spotType = Motorcycle
				case "A":
					spotType = Automobile
				default:
					active = false
				}

				grid[f][r][c] = &Spot{
					SpotType: spotType,
					IsActive: active,
				}
			}
		}
	}
	return &ParkingLot{
		floors:      floors,
		rows:        rows,
		cols:        cols,
		grid:        grid,
		vehicleMap:  make(map[string]string),
		lastSpotMap: make(map[string]string),
	}
}

func (p *ParkingLot) park(vehicleType VehicleType, vehicleNumber string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for f := 0; f < p.floors; f++ {
		for r := 0; r < p.rows; r++ {
			for c := 0; c < p.cols; c++ {
				spot := p.grid[f][r][c]
				if spot.IsActive && !spot.Occupied && spot.SpotType == vehicleType {
					spot.Occupied = true
					spot.VehicleNo = vehicleNumber

					spotId := fmt.Sprintf("%d-%d-%d", f, r, c)
					p.vehicleMap[vehicleNumber] = spotId
					p.lastSpotMap[vehicleNumber] = spotId
					return spotId, nil
				}
			}
		}
	}
	return "", errors.New("no available spot")
}

func (p *ParkingLot) unpark(spotId, vehicleNumber string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	loc := strings.Split(spotId, "-")
	if len(loc) != 3 {
		return errors.New("invalid spot id")
	}

	var f, r, c int
	_, _ = fmt.Sscanf(spotId, "%d-%d-%d", &f, &r, &c)

	if f >= p.floors || r >= p.rows || c >= p.cols {
		return errors.New("spot out of bounds")
	}

	spot := p.grid[f][r][c]
	if spot.VehicleNo != vehicleNumber || !spot.Occupied {
		return errors.New("vehicle not found in spot")
	}

	spot.Occupied = false
	spot.VehicleNo = ""
	delete(p.vehicleMap, vehicleNumber)
	return nil
}

func (p *ParkingLot) availableSpot(vehicleType VehicleType) []string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var spots []string
	for f := 0; f < p.floors; f++ {
		for r := 0; r < p.rows; r++ {
			for c := 0; c < p.cols; c++ {
				spot := p.grid[f][r][c]
				if spot.IsActive && !spot.Occupied && spot.SpotType == vehicleType {
					spots = append(spots, fmt.Sprintf("%d-%d-%d", f, r, c))
				}
			}
		}
	}
	return spots
}

func (p *ParkingLot) searchVehicle(vehicleNumber string) (string, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	spotId, found := p.lastSpotMap[vehicleNumber]
	return spotId, found
}

func main() {
	// Simulated layout with 1 floor, 2 rows, 3 columns
	layout := [][][]string{
		{
			{"B-1", "M-1", "X-0"},
			{"A-1", "M-1", "A-1"},
		},
	}
	pl := NewParkingLot(1, 2, 3, layout)

	spotID, err := pl.park(Bicycle, "BIKE123")
	fmt.Println("Park BIKE123:", spotID, err)

	spotID, err = pl.park(Motorcycle, "MOTO999")
	fmt.Println("Park MOTO999:", spotID, err)

	err = pl.unpark(spotID, "MOTO999")
	fmt.Println("Unpark MOTO999:", err)

	spots := pl.availableSpot(Motorcycle)
	fmt.Println("Available Motorcycles:", spots)

	lastSpot, found := pl.searchVehicle("MOTO999")
	fmt.Println("Last spot for MOTO999:", lastSpot, "Found:", found)
}
