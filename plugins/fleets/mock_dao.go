package fleets

import (
	"context"

	"gorm.io/gorm"

	"github.com/openshift-online/rh-trex-ai/pkg/errors"
)

var _ FleetDao = &fleetDaoMock{}

type fleetDaoMock struct {
	fleets FleetList
}

func NewMockFleetDao() *fleetDaoMock {
	return &fleetDaoMock{}
}

func (d *fleetDaoMock) Get(ctx context.Context, id string) (*Fleet, error) {
	for _, fleet := range d.fleets {
		if fleet.ID == id {
			return fleet, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (d *fleetDaoMock) Create(ctx context.Context, fleet *Fleet) (*Fleet, error) {
	d.fleets = append(d.fleets, fleet)
	return fleet, nil
}

func (d *fleetDaoMock) Replace(ctx context.Context, fleet *Fleet) (*Fleet, error) {
	return nil, errors.NotImplemented("Fleet").AsError()
}

func (d *fleetDaoMock) Delete(ctx context.Context, id string) error {
	return errors.NotImplemented("Fleet").AsError()
}

func (d *fleetDaoMock) FindByIDs(ctx context.Context, ids []string) (FleetList, error) {
	return nil, errors.NotImplemented("Fleet").AsError()
}

func (d *fleetDaoMock) All(ctx context.Context) (FleetList, error) {
	return d.fleets, nil
}
