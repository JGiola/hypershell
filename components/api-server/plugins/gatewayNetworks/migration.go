package gatewayNetworks

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/openshift-online/rh-trex-ai/pkg/db"
)

func migration() *gormigrate.Migration {
	type GatewayNetwork struct {
		db.Model
		Name         string
		FleetId      string
		Topology     *string
		TunnelMode   *string
		HubGatewayId *string
		Status       *string
	}

	return &gormigrate.Migration{
		ID: "2026080312548062",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&GatewayNetwork{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&GatewayNetwork{})
		},
	}
}
