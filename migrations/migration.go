package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/yodfhafx/go-crud/config"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1675671925CreateArticlesTable(),
			m1675745269CreateCategoriesTable(),
			m1675747773AddCategoryIDToArticles(),
			m1675752464CreateUsersTable(),
			m1675757483AddUserIDToArticles(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
