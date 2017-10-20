package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bouk/httprouter"
	"github.com/influxdata/chronograf"
)

type organizationRequest struct {
	Name string `json:"name"`
}

func (r *organizationRequest) ValidCreate() error {
	if r.Name == "" {
		return fmt.Errorf("Name required on Chronograf Organization request body")
	}
	return nil
}

func (r *organizationRequest) ValidUpdate() error {
	if r.Name == "" {
		return fmt.Errorf("No fields to update")
	}
	return nil
}

type organizationResponse struct {
	Links selfLinks `json:"links"`
	ID    uint64    `json:"id,string"`
	Name  string    `json:"name"`
}

func newOrganizationResponse(o *chronograf.Organization) *organizationResponse {
	return &organizationResponse{
		ID:   o.ID,
		Name: o.Name,
		Links: selfLinks{
			Self: fmt.Sprintf("/chronograf/v1/organizations/%d", o.ID),
		},
	}
}

type organizationsResponse struct {
	Links         selfLinks               `json:"links"`
	Organizations []*organizationResponse `json:"organizations"`
}

func newOrganizationsResponse(orgs []chronograf.Organization) *organizationsResponse {
	orgsResp := make([]*organizationResponse, len(orgs))
	for i, org := range orgs {
		orgsResp[i] = newOrganizationResponse(&org)
	}
	return &organizationsResponse{
		Organizations: orgsResp,
		Links: selfLinks{
			Self: "/chronograf/v1/organizations",
		},
	}
}

// Organizations retrieves all organizations from store
func (s *Service) Organizations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orgs, err := s.OrganizationsStore.All(ctx)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error(), s.Logger)
		return
	}

	res := newOrganizationsResponse(orgs)
	encodeJSON(w, http.StatusOK, res, s.Logger)
}

// NewOrganization adds a new organization to store
func (s *Service) NewOrganization(w http.ResponseWriter, r *http.Request) {
	var req organizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		invalidJSON(w, s.Logger)
		return
	}

	if err := req.ValidCreate(); err != nil {
		invalidData(w, err, s.Logger)
		return
	}

	ctx := r.Context()
	org := &chronograf.Organization{
		Name: req.Name,
	}

	res, err := s.OrganizationsStore.Add(ctx, org)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error(), s.Logger)
		return
	}

	co := newOrganizationResponse(res)
	location(w, co.Links.Self)
	encodeJSON(w, http.StatusCreated, co, s.Logger)
}

// UserID retrieves a organization with ID from store
func (s *Service) OrganizationID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := httprouter.GetParamFromContext(ctx, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("invalid organization id: %s", err.Error()), s.Logger)
		return
	}

	org, err := s.OrganizationsStore.Get(ctx, chronograf.OrganizationQuery{ID: &id})
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error(), s.Logger)
		return
	}

	res := newOrganizationResponse(org)
	encodeJSON(w, http.StatusOK, res, s.Logger)
}

func (s *Service) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	var req organizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		invalidJSON(w, s.Logger)
		return
	}

	if err := req.ValidUpdate(); err != nil {
		invalidData(w, err, s.Logger)
		return
	}

	ctx := r.Context()
	idStr := httprouter.GetParamFromContext(ctx, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("invalid organization id: %s", err.Error()), s.Logger)
		return
	}

	org, err := s.OrganizationsStore.Get(ctx, chronograf.OrganizationQuery{ID: &id})
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error(), s.Logger)
		return
	}

	if req.Name != "" {
		org.Name = req.Name
	}

	err = s.OrganizationsStore.Update(ctx, org)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error(), s.Logger)
		return
	}

	res := newOrganizationResponse(org)
	location(w, res.Links.Self)
	encodeJSON(w, http.StatusOK, res, s.Logger)

}
func (s *Service) RemoveOrganization(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := httprouter.GetParamFromContext(ctx, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("invalid organization id: %s", err.Error()), s.Logger)
		return
	}

	org, err := s.OrganizationsStore.Get(ctx, chronograf.OrganizationQuery{ID: &id})
	if err != nil {
		Error(w, http.StatusNotFound, err.Error(), s.Logger)
		return
	}
	if err := s.OrganizationsStore.Delete(ctx, org); err != nil {
		Error(w, http.StatusBadRequest, err.Error(), s.Logger)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
