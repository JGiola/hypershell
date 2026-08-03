package fleets_test

import (
	"context"
	"fmt"

	"github.com/openshift-online/hypershell/plugins/fleets"
	"github.com/openshift-online/rh-trex-ai/pkg/environments"
)

func newFleet(id string) (*fleets.Fleet, error) {
	fleetService := fleets.Service(&environments.Environment().Services)

	fleet := &fleets.Fleet{
		Name:        "test-name",
		Description: stringPtr("test-description"),
		Status:      stringPtr("test-status"),
	}

	sub, err := fleetService.Create(context.Background(), fleet)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func newFleetList(namePrefix string, count int) ([]*fleets.Fleet, error) {
	var items []*fleets.Fleet
	for i := 1; i <= count; i++ {
		name := fmt.Sprintf("%s_%d", namePrefix, i)
		c, err := newFleet(name)
		if err != nil {
			return nil, err
		}
		items = append(items, c)
	}
	return items, nil
}
func stringPtr(s string) *string { return &s }
