package skaioskit

import (
    "net/http"
    "encoding/json"
)

type ControllerResponse struct {
    Status int
    Body interface{}
}

//IController
type IController interface {
    //every controller should support these four methods... even if they just return 404
    Get(w http.ResponseWriter, r *http.Request) ControllerResponse
    Put(w http.ResponseWriter, r *http.Request) ControllerResponse
    Post(w http.ResponseWriter, r *http.Request) ControllerResponse
    Delete(w http.ResponseWriter, r *http.Request) ControllerResponse
}

//Controller wrapper to provide highlevel logic in a non dup'd way.
type ControllerProcessor struct {
    controller IController
}
func NewControllerProcessor(controller IController) *ControllerProcessor {
    return &ControllerProcessor{controller: controller}
}
func (p *ControllerProcessor) Logic(w http.ResponseWriter, r *http.Request) {
    resp := ControllerResponse{Status: http.StatusNotFound, Body: EmptyResponse{}}

    //route to the wrapped controller function based on request method
    switch r.Method {
        case "GET":
            resp = p.controller.Get(w, r)
            break
        case "POST":
            resp = p.controller.Post(w, r)
            break
        case "PUT":
            resp = p.controller.Put(w, r)
            break
        case "DELETE":
            resp = p.controller.Delete(w, r)
            break
    }

    p.writeJsonOutput(w, resp)
}
func (p *ControllerProcessor) writeJsonOutput(w http.ResponseWriter, resp ControllerResponse) {
    w.WriteHeader(resp.Status)
    jsonResp, err := json.Marshal(resp.Body);

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("{\"Success\": false,\"Message\": \"Cant Generate Json\"}"))
    } else {
        w.Write(jsonResp)
    }
}
