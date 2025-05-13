package parking

import (
	"context"
	"net/http"
	"time"

	"github.com/zuhrulumam/doit-test/business/entity"
	"github.com/zuhrulumam/doit-test/pkg"

	x "github.com/zuhrulumam/doit-test/pkg/errors"
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
		return result, x.WrapWithCode(err, http.StatusInternalServerError, "error get available parking spot")
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
		return x.WrapWithCode(err, http.StatusInternalServerError, "failed to insert vehicle")
	}

	return nil
}

func (p *parking) UpdateParkingSpot(ctx context.Context, data entity.UpdateParkingSpot) error {

	db := pkg.GetTransactionFromCtx(ctx, p.db)

	tx := db.WithContext(ctx).Model(&entity.ParkingSpot{})

	// Build conditional WHERE clause
	if data.ID > 0 {
		tx = tx.Where("id = ?", data.ID)
	} else if data.Floor > 0 && data.Row > 0 && data.Col > 0 {
		tx = tx.Where("floor = ? AND row = ? AND col = ?", data.Floor, data.Row, data.Col)
	} else {
		return x.NewWithCode(http.StatusBadRequest, "must provide either spot_id or (floor, row, col)")
	}

	updates := map[string]interface{}{}
	if data.Occupied != nil {
		updates["occupied"] = data.Occupied
	}

	if len(updates) == 0 {
		return x.NewWithCode(http.StatusBadRequest, "no updates provided")
	}

	if err := tx.Updates(updates).Error; err != nil {
		return x.WrapWithCode(err, http.StatusInternalServerError, "failed to update parking spot")
	}

	return nil
}

func (p *parking) UpdateVehicle(ctx context.Context, data entity.UpdateVehicle) error {
	db := pkg.GetTransactionFromCtx(ctx, p.db)

	if data.ID < 1 {
		return x.NewWithCode(http.StatusBadRequest, "vehicle id is required")
	}

	updates := map[string]interface{}{}
	if data.UnparkedAt != nil {
		updates["unparked_at"] = data.UnparkedAt
	}

	if len(updates) == 0 {
		return x.NewWithCode(http.StatusBadRequest, "no updates provided")
	}

	if err := db.WithContext(ctx).
		Model(&entity.Vehicle{}).
		Where("id = ?", data.ID).
		Updates(updates).Error; err != nil {
		return x.WrapWithCode(err, http.StatusInternalServerError, "failed to update vehicle")
	}

	return nil
}

func (p *parking) GetVehicle(ctx context.Context, data entity.SearchVehicle) (entity.Vehicle, error) {
	var (
		result entity.Vehicle
	)

	db := p.db.WithContext(ctx).Model(&entity.Vehicle{})

	// Filter by type
	if data.VehicleNumber != "" {
		db = db.Where("vehicle_number = ?", data.VehicleNumber)
	}

	// Only get the first match
	err := db.Order("id DESC").First(&result).Error
	if err != nil {
		return result, x.WrapWithCode(err, http.StatusInternalServerError, "failed get vehicle")
	}

	return result, nil
}
