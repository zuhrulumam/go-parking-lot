# 🅿️ Concurrent Multi-Floor Parking Lot System in Go

This project implements a **thread-safe**, **multi-floor**, and **concurrently accessible** parking lot system in Go. The parking lot supports different vehicle types, multiple entry/exit gates, and is designed for high-concurrency environments.

---

## 🚗 Supported Features

- ✅ Multiple floors (up to 8), rows (up to 1000), and columns (up to 1000)
- ✅ Active/inactive parking spots
- ✅ Supports 3 vehicle types:
  - Bicycles (`B-1`)
  - Motorcycles (`M-1`)
  - Automobiles (`A-1`)
- ✅ Thread-safe concurrent access using `sync.Mutex`
- ✅ Multi-gate simulation for parallel parking and unparking
- ✅ Full unit test coverage with table-driven tests
- ✅ Stress test with 1000+ goroutines for benchmarking concurrency

---

## 🏗️ System Design

### Data Model

- **ParkingLot**: Manages all floors and spots.
- **Spot**: Represents an individual parking spot.
- **Gate**: Simulates concurrent access (via goroutines).
- **Vehicle Mapping**: Tracks which vehicle is parked at which spot, even after it is unparked.

---

## 🔧 Key Functions

| Function                           | Description                                        |
| ---------------------------------- | -------------------------------------------------- |
| `park(vehicleType, vehicleNumber)` | Parks a vehicle and returns assigned spot ID       |
| `unpark(spotId, vehicleNumber)`    | Unparks a vehicle from the given spot              |
| `availableSpot(vehicleType)`       | Lists all available spots for a given vehicle type |
| `searchVehicle(vehicleNumber)`     | Returns last known spot for the vehicle            |

---

## 🧪 Testing

### Unit Tests

We use table-driven unit tests to verify each function with multiple scenarios using `for` loops.

Run tests:

```bash
go test -v
```
