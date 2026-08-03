package gateways

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/openshift-online/rh-trex-ai/pkg/db"
)

func migration() *gormigrate.Migration {
	type Gateway struct {
		db.Model
		Name        string
		FleetId     string
		ClusterId   string
		ReleaseId   string
		DatabaseId  string
		Namespace   string
		ExternalDns *string
		TlsMode     *string
		ServiceType *string
		Status      *string
		Phase       *string
	}

	return &gormigrate.Migration{
		ID: "2026080312546877",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Gateway{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Gateway{})
		},
	}
}
