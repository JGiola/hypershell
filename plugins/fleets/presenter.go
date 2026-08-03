package fleets

import (
	"github.com/openshift-online/hypershell/pkg/api/openapi"
	"github.com/openshift-online/rh-trex-ai/pkg/api"
	"github.com/openshift-online/rh-trex-ai/pkg/api/presenters"
	"github.com/openshift-online/rh-trex-ai/pkg/util"
)

func ConvertFleet(fleet openapi.Fleet) *Fleet {
	c := &Fleet{
		Meta: api.Meta{
			ID: util.NilToEmptyString(fleet.Id),
		},
	}
	c.Name = fleet.Name
	c.Description = fleet.Description
	c.Status = fleet.Status

	if fleet.CreatedAt != nil {
		c.CreatedAt = *fleet.CreatedAt
		c.UpdatedAt = *fleet.UpdatedAt
	}

	return c
}

func PresentFleet(fleet *Fleet) openapi.Fleet {
	reference := presenters.PresentReference(fleet.ID, fleet)
	return openapi.Fleet{
		Id:          reference.Id,
		Kind:        reference.Kind,
		Href:        reference.Href,
		CreatedAt:   openapi.PtrTime(fleet.CreatedAt),
		UpdatedAt:   openapi.PtrTime(fleet.UpdatedAt),
		Name:        fleet.Name,
		Description: fleet.Description,
		Status:      fleet.Status,
	}
}
