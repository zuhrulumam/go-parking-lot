package pkg

import (
	"context"

	"gorm.io/gorm"
)

func GetTransactionFromCtx(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(TxCtxValue).(*gorm.DB)
	if !ok {
		return db
	}

	return tx
}
