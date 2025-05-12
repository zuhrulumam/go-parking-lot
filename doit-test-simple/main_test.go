package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func createTestLot() *ParkingLot {
	layout := [][][]string{
		{
			{"B-1", "M-1", "X-0"},
			{"A-1", "M-1", "A-1"},
		},
	}
	return NewParkingLot(1, 2, 3, layout)
}

func TestPark(t *testing.T) {
	testCases := []struct {
		vehicleType    VehicleType
		vehicleNumber  string
		expectError    bool
		expectedSpotID string
	}{
		{Bicycle, "B-001", false, "0-0-0"},
		{Motorcycle, "M-001", false, "0-0-1"},
		{Motorcycle, "M-002", false, "0-1-1"},
		{Motorcycle, "M-003", true, ""},
		{Automobile, "A-001", false, "0-1-0"},
		{Automobile, "A-002", false, "0-1-2"},
		{Automobile, "A-003", true, ""},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Park[%d] %s", i, tc.vehicleNumber), func(t *testing.T) {
			pl := createTestLot()

			// Fill up for previous test cases
			for j := 0; j < i; j++ {
				_, _ = pl.park(testCases[j].vehicleType, testCases[j].vehicleNumber)
			}

			spotID, err := pl.park(tc.vehicleType, tc.vehicleNumber)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got spotID: %s", spotID)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tc.expectedSpotID != "" && spotID != tc.expectedSpotID {
					t.Errorf("Expected spotID %s, got %s", tc.expectedSpotID, spotID)
				}
			}
		})
	}
}

func TestUnpark(t *testing.T) {
	testCases := []struct {
		name           string
		vehicleType    VehicleType
		vehicleNumber  string
		modifyBefore   func(*ParkingLot, string) // optional modification before test
		spotIDExpected string
		expectError    bool
	}{
		{
			name:          "Valid unpark",
			vehicleType:   Automobile,
			vehicleNumber: "CAR1",
			modifyBefore:  nil,
			expectError:   false,
		},
		{
			name:          "Unpark twice",
			vehicleType:   Motorcycle,
			vehicleNumber: "MOTO2",
			modifyBefore: func(pl *ParkingLot, spot string) {
				_ = pl.unpark(spot, "MOTO2")
			},
			expectError: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Unpark[%d] %s", i, tc.name), func(t *testing.T) {
			pl := createTestLot()
			spotID, _ := pl.park(tc.vehicleType, tc.vehicleNumber)

			if tc.modifyBefore != nil {
				tc.modifyBefore(pl, spotID)
			}

			err := pl.unpark(spotID, tc.vehicleNumber)
			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestAvailableSpots(t *testing.T) {
	vehicleTypes := []VehicleType{Bicycle, Motorcycle, Automobile}

	for _, vt := range vehicleTypes {
		t.Run(fmt.Sprintf("Available before and after park [%s]", vt), func(t *testing.T) {
			pl := createTestLot()

			initial := pl.availableSpot(vt)
			if len(initial) == 0 {
				t.Errorf("Expected available spots for %s", vt)
			}

			_, _ = pl.park(vt, "TEST-AVAIL")
			after := pl.availableSpot(vt)

			if len(after) != len(initial)-1 {
				t.Errorf("Expected %d spots after parking, got %d", len(initial)-1, len(after))
			}
		})
	}
}

func TestSearchVehicle(t *testing.T) {
	testCases := []struct {
		name          string
		vehicleType   VehicleType
		vehicleNumber string
		unparkFirst   bool
		expectFound   bool
	}{
		{"Search while parked", Bicycle, "B-123", false, true},
		{"Search after unpark", Automobile, "A-321", true, true},
		{"Search unknown vehicle", Motorcycle, "X-999", false, false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Search[%d] %s", i, tc.name), func(t *testing.T) {
			pl := createTestLot()

			var spotID string
			if tc.expectFound {
				spotID, _ = pl.park(tc.vehicleType, tc.vehicleNumber)
				if tc.unparkFirst {
					_ = pl.unpark(spotID, tc.vehicleNumber)
				}
			}

			spot, found := pl.searchVehicle(tc.vehicleNumber)
			if found != tc.expectFound {
				t.Errorf("Expected found=%v, got %v", tc.expectFound, found)
			}
			if found && spot != spotID {
				t.Errorf("Expected spotID %s, got %s", spotID, spot)
			}
		})
	}
}

// test concurency
type Gate struct {
	id  int
	lot *ParkingLot
}

func TestConcurrentParkingFromMultipleGates(t *testing.T) {
	pl := createTestLot()

	var wg sync.WaitGroup
	gateCount := 3
	vehiclePrefix := []string{"A", "B", "C"}

	for i := 0; i < gateCount; i++ {
		wg.Add(1)
		go func(gateID int) {
			defer wg.Done()
			gate := Gate{id: gateID, lot: pl}
			for j := 0; j < 3; j++ { // each gate parks 3 vehicles
				vehicleNum := fmt.Sprintf("%s-%d", vehiclePrefix[gateID], j)
				_, err := gate.lot.park(Automobile, vehicleNum)
				if err != nil {
					t.Logf("Gate %d failed to park %s: %v", gateID, vehicleNum, err)
				} else {
					t.Logf("Gate %d successfully parked %s", gateID, vehicleNum)
				}
			}
		}(i)
	}
	wg.Wait()

	total := len(pl.availableSpot(Automobile))
	t.Logf("Remaining available automobile spots: %d", total)
}

func TestStressConcurrentGates(t *testing.T) {
	// Create a large parking lot for stress testing
	floors, rows, cols := 2, 50, 50 // Total spots: 2 * 50 * 50 = 5000
	layout := make([][][]string, floors)
	for f := 0; f < floors; f++ {
		layout[f] = make([][]string, rows)
		for r := 0; r < rows; r++ {
			layout[f][r] = make([]string, cols)
			for c := 0; c < cols; c++ {
				layout[f][r][c] = "A-1" // All automobile active spots
			}
		}
	}

	pl := NewParkingLot(floors, rows, cols, layout)

	var wg sync.WaitGroup
	totalGoroutines := 1000
	errors := int32(0)

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			vehicleNumber := fmt.Sprintf("CAR-%d", id)

			// Try to park
			spot, err := pl.park(Automobile, vehicleNumber)
			if err != nil {
				atomic.AddInt32(&errors, 1)
				return
			}

			// Simulate some usage (could be replaced with time.Sleep)
			// time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))

			// Unpark
			err = pl.unpark(spot, vehicleNumber)
			if err != nil {
				atomic.AddInt32(&errors, 1)
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Finished %d concurrent park/unpark operations", totalGoroutines)
	t.Logf("Total errors during test: %d", errors)
	if errors > 0 {
		t.Errorf("Some operations failed under concurrency")
	}
}
