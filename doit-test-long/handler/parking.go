package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zuhrulumam/doit-test/business/entity"

	x "github.com/zuhrulumam/doit-test/pkg/errors"
)

var (
	validate = validator.New()
)

// SearchVehicle godoc
// @Summary      Search a parked vehicle
// @Description  Returns information about a vehicle parked in the lot
// @Tags         Parking
// @Accept       json
// @Produce      json
// @Param        vehicle_number query string true "Vehicle Number"
// @Success      200 {object} handler.SearchVehicleResponse
// @Failure      400 {object} handler.ErrorResponse
// @Router       /vehicle/search [get]
func (e *rest) SearchVehicle(c *fiber.Ctx) error {

	vehicleNumber := c.Query("vehicle_number")

	veh, err := e.uc.Parking.SearchVehicle(c.Context(), entity.SearchVehicle{
		VehicleNumber: vehicleNumber,
	})
	if err != nil {
		return e.compileError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(SearchVehicleResponse{
		Success: true,
		Message: "Done Search vehicle !",
		Vehicle: &veh,
	})
}

// AvailableSpot godoc
// @Summary      Get available parking spots
// @Description  Returns a list of available spots for a specific vehicle type
// @Tags         Parking
// @Accept       json
// @Produce      json
// @Param        vehicle_type query string true "Vehicle Type (M, B, A)"
// @Success      200 {object} handler.AvailableSpotResponse
// @Failure      400 {object} handler.ErrorResponse
// @Router       /spot/available [get]
func (e *rest) AvailableSpot(c *fiber.Ctx) error {

	var (
		res         []ParkingSpotBrief
		vehicleType = c.Query("vehicle_type")
	)

	spots, err := e.uc.Parking.AvailableSpot(c.Context(), entity.GetAvailablePark{
		VehicleType: entity.VehicleType(vehicleType),
	})
	if err != nil {
		return e.compileError(c, err)
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

// Park godoc
// @Summary      Park a vehicle
// @Description  Parks a vehicle into an available spot
// @Tags         Parking
// @Accept       json
// @Produce      json
// @Param        body body handler.ParkRequest true "Vehicle Info"
// @Success      200 {object} handler.ParkResponse
// @Failure      400 {object} handler.ErrorResponse
// @Router       /vehicle/park [post]
func (e *rest) Park(c *fiber.Ctx) error {

	var input ParkRequest
	if err := c.BodyParser(&input); err != nil {
		return e.compileError(c, x.WrapWithCode(err, http.StatusBadRequest, "invalid input"))
	}

	if err := validate.Struct(input); err != nil {
		return e.compileError(c, x.WrapWithCode(err, http.StatusBadRequest, "failed validation"))
	}

	err := e.uc.Parking.Park(c.Context(), entity.Park{
		VehicleType:   entity.VehicleType(input.VehicleType),
		VehicleNumber: input.VehicleNumber,
	})
	if err != nil {
		return e.compileError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(ParkResponse{
		Success: true,
		Message: "Done parking vehicle !",
	})
}

// UnPark godoc
// @Summary      Unpark a vehicle
// @Description  Removes a vehicle from the parking lot
// @Tags         Parking
// @Accept       json
// @Produce      json
// @Param        body body handler.UnparkRequest true "Unpark Info"
// @Success      200 {object} handler.UnparkResponse
// @Failure      400 {object} handler.ErrorResponse
// @Router       /vehicle/unpark [post]
func (e *rest) UnPark(c *fiber.Ctx) error {

	var input UnparkRequest
	if err := c.BodyParser(&input); err != nil {
		return e.compileError(c, x.WrapWithCode(err, http.StatusBadRequest, "invalid input"))
	}

	if err := validate.Struct(input); err != nil {
		return e.compileError(c, x.WrapWithCode(err, http.StatusBadRequest, "failed validation"))
	}

	err := e.uc.Parking.Unpark(c.Context(), entity.UnPark{
		SpotID:        input.SpotID,
		VehicleNumber: input.VehicleNumber,
	})
	if err != nil {
		return e.compileError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(UnparkResponse{
		Success: true,
		Message: "Done unparking vehicle !",
	})
}
