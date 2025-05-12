package parking

import (
	"context"

	"github.com/zuhrulumam/doit-test/business/entity"
)

func (p *parking) Park(ctx context.Context, data entity.Park) error {

	p.TransactionDom.RunInTx(ctx, func(newCtx context.Context) error {

		return nil
	})

	// check parking_spot by vehicle type and active

	// insert vehicle

	// update parking_spot active = false as floor row col

	return nil
}

func (p *parking) Unpark(ctx context.Context, data entity.UnPark) error {

	// get vehicle by spotID, vehicle number, and UnparkedAt null

	// update vehicle

	// update parking_spot to active = true

	return nil
}

func (p *parking) AvailableSpot(ctx context.Context, data entity.GetAvailablePark) {
	// check parking_spot by vehicle type and active

	// return available
}

func (p *parking) SearchVehicle(ctx context.Context, data entity.SearchVehicle) {
	// get vehicle number

	// return park
}
