package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/yodfhafx/go-crud/models"
	"gorm.io/gorm"
)

func m1675757483AddUserIDToArticles() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1675757483",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Article{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn(&models.Article{}, "user_id")
		},
	}
}
