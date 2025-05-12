package pkg

import (
	"context"

	"gorm.io/gorm"
)

func GetTransactionFromCtx(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		return db
	}

	return tx
}
