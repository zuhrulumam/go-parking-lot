package pkg

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zuhrulumam/go-parking-lot/business/entity"
)

type CtxVal string

var (
	TxCtxValue CtxVal = "tx"
)

func BoolPtr(b bool) *bool {
	return &b
}

func TimePtr(b time.Time) *time.Time {
	return &b
}

func ParseSpotID(spotID string) (*entity.SpotID, error) {
	parts := strings.Split(spotID, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid spotID format")
	}

	floor, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid floor: %w", err)
	}

	row, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid col: %w", err)
	}

	col, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid col: %w", err)
	}

	return &entity.SpotID{
		Floor: floor,
		Row:   row,
		Col:   col,
	}, nil
}
