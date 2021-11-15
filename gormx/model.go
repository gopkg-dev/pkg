package gormx

import (
	"context"
	"time"

	"github.com/gopkg-dev/pkg/contextx"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

// Model base model
type Model struct {
	ID        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at;index;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// GetDB Get gorm.DB from context
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FromTrans(ctx)
	if ok && !contextx.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextx.FromTransLock(ctx) {
				db = db.Clauses(clause.Locking{Strength: "UPDATE"})
			}
			return db
		}
	}
	return defDB
}

// GetDBWithModel Get gorm.DB.Model from context
func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(m)
}
