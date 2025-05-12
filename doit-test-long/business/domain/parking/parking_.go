package parking

import (
	"context"
	"fmt"
	"time"

	"github.com/zuhrulumam/doit-test/business/entity"
	"github.com/zuhrulumam/doit-test/pkg"
)

func (p *parking) GetAvailableParkingSpot(ctx context.Context, data entity.GetAvailableParkingSpot) ([]entity.ParkingSpot, error) {

	var (
		result []entity.ParkingSpot
	)

	db := p.db.WithContext(ctx).Model(&entity.ParkingSpot{})

	// Filter by type
	if data.VehicleType != "" {
		db = db.Where("type = ?", data.VehicleType)
	}

	// Filter by active status
	if data.Active != nil {
		db = db.Where("active = ?", *data.Active)
	}

	// Filter by occupied status
	if data.Occupied != nil {
		db = db.Where("occupied = ?", *data.Occupied)
	}

	// Only get the first match
	err := db.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *parking) InsertVehicle(ctx context.Context, data entity.InsertVehicle) error {
	db := pkg.GetTransactionFromCtx(ctx, p.db)

	vehicle := entity.Vehicle{
		VehicleNumber: data.VehicleNumber,
		VehicleType:   data.VehicleType,
		SpotID:        data.SpotID,
		ParkedAt:      time.Now(),
	}

	if err := db.WithContext(ctx).Create(&vehicle).Error; err != nil {
		return fmt.Errorf("failed to insert vehicle: %w", err)
	}

	return nil
}

func (p *parking) UpdateParkingSpot(ctx context.Context, data entity.UpdateParkingSpot) error {

	db := pkg.GetTransactionFromCtx(ctx, p.db)

	// if data.SpotID == "" {
	// 	return fmt.Errorf("spot_id is required")
	// }

	if err := db.WithContext(ctx).
		Model(&entity.ParkingSpot{}).
		// Where("spot_id = ?", data.SpotID).
		Update("occupied", data.Occupied).Error; err != nil {
		return fmt.Errorf("failed to update parking spot: %w", err)
	}

	return nil
}

func (p *parking) UpdateVehicle(ctx context.Context, data entity.UpdateVehicle) error {
	db := pkg.GetTransactionFromCtx(ctx, p.db)

	if data.VehicleNumber == "" {
		return fmt.Errorf("vehicle_number is required")
	}

	updates := map[string]interface{}{}
	if data.UnparkedAt != nil {
		updates["unparked_at"] = data.UnparkedAt
	}

	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	if err := db.WithContext(ctx).
		Model(&entity.Vehicle{}).
		Where("vehicle_number = ?", data.VehicleNumber).
		Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update vehicle: %w", err)
	}

	return nil
}
