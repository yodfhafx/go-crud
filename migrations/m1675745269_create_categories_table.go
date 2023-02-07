package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/yodfhafx/go-crud/models"
	"gorm.io/gorm"
)

func m1675745269CreateCategoriesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1675745269",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Category{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("categories")
		},
	}
}
