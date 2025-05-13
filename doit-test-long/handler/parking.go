package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zuhrulumam/doit-test/business/entity"
)

func (e *rest) SearchVehicle(c *fiber.Ctx) error {

	vehicleNumber := c.Query("vehicle_number")

	veh, err := e.uc.Parking.SearchVehicle(c.Context(), entity.SearchVehicle{
		VehicleNumber: vehicleNumber,
	})
	if err != nil {
		c.Status(400)
	}

	return c.Status(fiber.StatusOK).JSON(SearchVehicleResponse{
		Success: true,
		Message: "Done Search vehicle !",
		Vehicle: &veh,
	})
}

func (e *rest) AvailableSpot(c *fiber.Ctx) error {

	var (
		res         []ParkingSpotBrief
		vehicleType = c.Query("vehicle_type")
	)

	spots, err := e.uc.Parking.AvailableSpot(c.Context(), entity.GetAvailablePark{
		VehicleType: entity.VehicleType(vehicleType),
	})
	if err != nil {
		c.Status(400)
	}

	for _, s := range spots {
		res = append(res, ParkingSpotBrief{
			SpotID: fmt.Sprintf("%d-%d-%d", s.Floor, s.Row, s.Col),
			Floor:  s.Floor,
			Row:    s.Row,
			Column: s.Col,
		})
	}

	return c.Status(fiber.StatusOK).JSON(AvailableSpotResponse{
		Success:        true,
		Message:        "Done Search vehicle !",
		AvailableSpots: res,
		VehicleType:    vehicleType,
	})
}

func (e *rest) Park(c *fiber.Ctx) error {

	var input ParkRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	err := e.uc.Parking.Park(c.Context(), entity.Park{
		VehicleType:   entity.VehicleType(input.VehicleType),
		VehicleNumber: input.VehicleNumber,
	})
	if err != nil {
		c.Status(400)
	}

	return c.Status(fiber.StatusOK).JSON(ParkResponse{
		Success: true,
		Message: "Done parking vehicle !",
	})
}

func (e *rest) UnPark(c *fiber.Ctx) error {

	var input UnparkRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	err := e.uc.Parking.Unpark(c.Context(), entity.UnPark{
		SpotID:        input.SpotID,
		VehicleNumber: input.VehicleNumber,
	})
	if err != nil {
		c.Status(400)
	}

	return c.Status(fiber.StatusOK).JSON(UnparkResponse{
		Success: true,
		Message: "Done unparking vehicle !",
	})
}
