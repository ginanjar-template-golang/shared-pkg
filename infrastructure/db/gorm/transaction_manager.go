package gorm

import (
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"gorm.io/gorm"
)

type TransactionManager struct {
	DB *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{DB: db}
}

// TxFunc callback menerima semua repository yang dibungkus
type TxFunc func(repos map[string]any) error

// Transaction jalankan callback dalam satu transaksi
func (tm *TransactionManager) Transaction(
	fn TxFunc,
	repoConstructors map[string]func(tx *gorm.DB) any,
) error {
	tx := tm.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	repos := make(map[string]any)
	for key, constructor := range repoConstructors {
		repos[key] = constructor(tx)
	}

	if err := fn(repos); err != nil {
		tx.Rollback()
		logger.Error("Transaction rollback due to error", map[string]any{"error": err})
		return err
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error("Transaction commit failed", map[string]any{"error": err})
		return err
	}

	return nil
}
