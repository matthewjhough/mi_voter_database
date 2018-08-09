package controllers

import (
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/core"
    "skaioskit/services"
)

type VoterController struct {
    voterService services.IVoterService
}
func NewVoterController(voterService services.IVoterService) *VoterController {
    return &VoterController{
        voterService: voterService,
    }
}
func (p *VoterController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    first := r.URL.Query().Get("first")
    last := r.URL.Query().Get("last")

    //voters, err := p.voterService.GetVoters(core.QueryRequest{})
    voters, err := p.voterService.GetVotersByName(first, last)
    if err == nil {
        return skaioskit.ControllerResponse{Status: http.StatusOK, Body: core.GetVotersResponse{Voters: voters}}
    } else {
        return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
    }
}
func (p *VoterController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *VoterController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    count := p.voterService.GetVoterCount()
    return skaioskit.ControllerResponse{Status: http.StatusOK, Body: core.GetVoterCount{Count: count}}
}
func (p *VoterController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
