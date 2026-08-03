package fleets

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/openshift-online/hypershell/pkg/api/openapi"
	"github.com/openshift-online/rh-trex-ai/pkg/api/presenters"
	"github.com/openshift-online/rh-trex-ai/pkg/errors"
	"github.com/openshift-online/rh-trex-ai/pkg/handlers"
	"github.com/openshift-online/rh-trex-ai/pkg/services"
)

var _ handlers.RestHandler = fleetHandler{}

type fleetHandler struct {
	fleet   FleetService
	generic services.GenericService
}

func NewFleetHandler(fleet FleetService, generic services.GenericService) *fleetHandler {
	return &fleetHandler{
		fleet:   fleet,
		generic: generic,
	}
}

func (h fleetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var fleet openapi.Fleet
	cfg := &handlers.HandlerConfig{
		Body: &fleet,
		Validators: []handlers.Validate{
			handlers.ValidateEmpty(&fleet, "Id", "id"),
		},
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			fleetModel := ConvertFleet(fleet)
			fleetModel, err := h.fleet.Create(ctx, fleetModel)
			if err != nil {
				return nil, err
			}
			return PresentFleet(fleetModel), nil
		},
		ErrorHandler: handlers.HandleError,
	}

	handlers.Handle(w, r, cfg, http.StatusCreated)
}

func (h fleetHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var patch openapi.FleetPatchRequest

	cfg := &handlers.HandlerConfig{
		Body:       &patch,
		Validators: []handlers.Validate{},
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			id := mux.Vars(r)["id"]
			found, err := h.fleet.Get(ctx, id)
			if err != nil {
				return nil, err
			}

			if patch.Name != nil {
				found.Name = *patch.Name
			}
			if patch.Description != nil {
				found.Description = patch.Description
			}
			if patch.Status != nil {
				found.Status = patch.Status
			}

			fleetModel, err := h.fleet.Replace(ctx, found)
			if err != nil {
				return nil, err
			}
			return PresentFleet(fleetModel), nil
		},
		ErrorHandler: handlers.HandleError,
	}

	handlers.Handle(w, r, cfg, http.StatusOK)
}

func (h fleetHandler) List(w http.ResponseWriter, r *http.Request) {
	cfg := &handlers.HandlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()

			listArgs := services.NewListArguments(r.URL.Query())
			var fleets []Fleet
			paging, err := h.generic.List(ctx, "id", listArgs, &fleets)
			if err != nil {
				return nil, err
			}
			kindStr := "FleetList"
			pageVal := int32(paging.Page)
			sizeVal := int32(paging.Size)
			totalVal := int32(paging.Total)
			fleetList := openapi.FleetList{
				Kind:  &kindStr,
				Page:  &pageVal,
				Size:  &sizeVal,
				Total: &totalVal,
				Items: []openapi.Fleet{},
			}

			for _, fleet := range fleets {
				converted := PresentFleet(&fleet)
				fleetList.Items = append(fleetList.Items, converted)
			}
			if listArgs.Fields != nil {
				filteredItems, err := presenters.SliceFilter(listArgs.Fields, fleetList.Items)
				if err != nil {
					return nil, err
				}
				return filteredItems, nil
			}
			return fleetList, nil
		},
	}

	handlers.HandleList(w, r, cfg)
}

func (h fleetHandler) Get(w http.ResponseWriter, r *http.Request) {
	cfg := &handlers.HandlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			ctx := r.Context()
			fleet, err := h.fleet.Get(ctx, id)
			if err != nil {
				return nil, err
			}

			return PresentFleet(fleet), nil
		},
	}

	handlers.HandleGet(w, r, cfg)
}

func (h fleetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	cfg := &handlers.HandlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			ctx := r.Context()
			err := h.fleet.Delete(ctx, id)
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
	}
	handlers.HandleDelete(w, r, cfg, http.StatusNoContent)
}
