package managedDatabases

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/openshift-online/rh-trex-ai/pkg/db"
)

func migration() *gormigrate.Migration {
	type ManagedDatabase struct {
		db.Model
		Name             string
		FleetId          string
		Provider         string
		Region           *string
		Engine           *string
		EngineVersion    *string
		InstanceClass    *string
		ConnectionSecret *string
		Status           *string
	}

	return &gormigrate.Migration{
		ID: "2026080312542176",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&ManagedDatabase{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&ManagedDatabase{})
		},
	}
}
