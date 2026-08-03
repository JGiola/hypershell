package fleets

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/openshift-online/rh-trex-ai/pkg/api"
	"github.com/openshift-online/rh-trex-ai/pkg/db"
)

type FleetDao interface {
	Get(ctx context.Context, id string) (*Fleet, error)
	Create(ctx context.Context, fleet *Fleet) (*Fleet, error)
	Replace(ctx context.Context, fleet *Fleet) (*Fleet, error)
	Delete(ctx context.Context, id string) error
	FindByIDs(ctx context.Context, ids []string) (FleetList, error)
	All(ctx context.Context) (FleetList, error)
}

var _ FleetDao = &sqlFleetDao{}

type sqlFleetDao struct {
	sessionFactory *db.SessionFactory
}

func NewFleetDao(sessionFactory *db.SessionFactory) FleetDao {
	return &sqlFleetDao{sessionFactory: sessionFactory}
}

func (d *sqlFleetDao) Get(ctx context.Context, id string) (*Fleet, error) {
	g2 := (*d.sessionFactory).New(ctx)
	var fleet Fleet
	if err := g2.Take(&fleet, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &fleet, nil
}

func (d *sqlFleetDao) Create(ctx context.Context, fleet *Fleet) (*Fleet, error) {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Create(fleet).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, err
	}
	return fleet, nil
}

func (d *sqlFleetDao) Replace(ctx context.Context, fleet *Fleet) (*Fleet, error) {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Save(fleet).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, err
	}
	return fleet, nil
}

func (d *sqlFleetDao) Delete(ctx context.Context, id string) error {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Delete(&Fleet{Meta: api.Meta{ID: id}}).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return err
	}
	return nil
}

func (d *sqlFleetDao) FindByIDs(ctx context.Context, ids []string) (FleetList, error) {
	g2 := (*d.sessionFactory).New(ctx)
	fleets := FleetList{}
	if err := g2.Where("id in (?)", ids).Find(&fleets).Error; err != nil {
		return nil, err
	}
	return fleets, nil
}

func (d *sqlFleetDao) All(ctx context.Context) (FleetList, error) {
	g2 := (*d.sessionFactory).New(ctx)
	fleets := FleetList{}
	if err := g2.Find(&fleets).Error; err != nil {
		return nil, err
	}
	return fleets, nil
}
