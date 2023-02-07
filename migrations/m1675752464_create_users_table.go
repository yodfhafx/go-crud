package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/yodfhafx/go-crud/models"
	"gorm.io/gorm"
)

func m1675752464CreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1675752464",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
