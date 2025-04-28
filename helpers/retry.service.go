package helpers

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type DBWithRetry struct {
	db         *gorm.DB
	maxRetries int
}

func NewDBWithRetry(db *gorm.DB) *DBWithRetry {
	return &DBWithRetry{
		db:         db,
		maxRetries: 5, // increased retries for better reliability
	}
}

func (dbr *DBWithRetry) CreateWithDynamicGenerator(value interface{}, regenerateFunc func() error) error {
	var err error
	for attempts := 0; attempts < dbr.maxRetries; attempts++ {
		tx := dbr.db.Begin()
		if tx.Error != nil {
			return tx.Error
		}

		err = tx.Create(value).Error
		if err == nil {
			return tx.Commit().Error
		}

		tx.Rollback()

		if isDeadlockError(err) {
			// Deadlock detected, retry directly
			time.Sleep(100 * time.Millisecond)
			continue
		} else if isDuplicateKeyError(err) {
			// Duplicate detected, regenerate code
			if regenerateFunc != nil {
				if regenErr := regenerateFunc(); regenErr != nil {
					return fmt.Errorf("regeneration failed: %w", regenErr)
				}
			}
			time.Sleep(50 * time.Millisecond)
			continue
		} else {
			// Some other error
			return err
		}
	}
	return fmt.Errorf("maximum retries exceeded: %w", err)
}

func isDeadlockError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "deadlock") || strings.Contains(errMsg, "lock wait timeout")
}

func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "duplicate entry")
}
