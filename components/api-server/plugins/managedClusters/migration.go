package managedClusters

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/openshift-online/rh-trex-ai/pkg/db"
)

func migration() *gormigrate.Migration {
	type ManagedCluster struct {
		db.Model
		Name             string
		FleetId          string
		Provider         string
		Region           *string
		KubeconfigSecret string
		Status           *string
		ApiServerUrl     *string
	}

	return &gormigrate.Migration{
		ID: "2026080312543137",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&ManagedCluster{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&ManagedCluster{})
		},
	}
}
