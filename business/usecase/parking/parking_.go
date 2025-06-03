package parking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/zuhrulumam/go-parking-lot/business/entity"
	"github.com/zuhrulumam/go-parking-lot/pkg"
	x "github.com/zuhrulumam/go-parking-lot/pkg/errors"
)

func (p *parking) Park(ctx context.Context, data entity.Park) error {

	err := p.TransactionDom.RunInTx(ctx, func(newCtx context.Context) error {

		// check parking_spot by vehicle type, active, and not occupied
		pSpots, err := p.ParkingDom.GetAvailableParkingSpot(newCtx, entity.GetAvailableParkingSpot{
			VehicleType: data.VehicleType,
			Active:      pkg.BoolPtr(true),
			Occupied:    pkg.BoolPtr(false),
			UseLock:     true,
		})
		if err != nil {
			return err
		}

		if len(pSpots) < 1 {
			return errors.New("no available parking")
		}

		spot := pSpots[0]
		spotID := fmt.Sprintf("%d-%d-%d", spot.Floor, spot.Row, spot.Col)

		// update parking_spot occupied = true as floor row col
		err = p.ParkingDom.UpdateParkingSpot(newCtx, entity.UpdateParkingSpot{
			ID:       spot.ID,
			Occupied: pkg.BoolPtr(true),
		})
		if err != nil {
			return err
		}

		// insert vehicle
		err = p.ParkingDom.InsertVehicle(newCtx, entity.InsertVehicle{
			VehicleNumber: data.VehicleNumber,
			VehicleType:   string(data.VehicleType),
			SpotID:        spotID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *parking) Unpark(ctx context.Context, data entity.UnPark) error {

	return p.TransactionDom.RunInTx(ctx, func(newCtx context.Context) error {

		// get vehicle by spotID, vehicle number, and UnparkedAt null
		vec, err := p.ParkingDom.GetVehicle(newCtx, entity.SearchVehicle{
			VehicleNumber: data.VehicleNumber,
		})
		if err != nil {
			return err
		}

		if vec.UnparkedAt != nil {
			return x.NewWithCode(http.StatusBadRequest, "already unparked")
		}

		// update vehicle
		err = p.ParkingDom.UpdateVehicle(newCtx, entity.UpdateVehicle{
			ID:         vec.ID,
			UnparkedAt: pkg.TimePtr(time.Now()),
		})
		if err != nil {
			return err
		}

		sp, err := pkg.ParseSpotID(vec.SpotID)
		if err != nil {
			return err
		}

		// update parking_spot to occupied = false
		err = p.ParkingDom.UpdateParkingSpot(newCtx, entity.UpdateParkingSpot{
			Floor:    sp.Floor,
			Row:      sp.Row,
			Col:      sp.Col,
			Occupied: pkg.BoolPtr(false),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (p *parking) AvailableSpot(ctx context.Context, data entity.GetAvailablePark) ([]entity.ParkingSpot, error) {
	// check parking_spot by vehicle type, active, and not occupied
	return p.ParkingDom.GetAvailableParkingSpot(ctx, entity.GetAvailableParkingSpot{
		VehicleType: data.VehicleType,
		Active:      pkg.BoolPtr(true),
		Occupied:    pkg.BoolPtr(false),
	})

}

func (p *parking) SearchVehicle(ctx context.Context, data entity.SearchVehicle) (entity.Vehicle, error) {
	return p.ParkingDom.GetVehicle(ctx, data)
}
