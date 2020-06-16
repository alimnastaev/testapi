package router

import (
	"encoding/json"
	"fmt"
	"github.com/alimnastaev/testapi/internal/api/controller"
	"net/http"
)

// Router is going to handle the requests
type Router struct {
	controllerSvc controller.API
	version       string
}

// New returns a new implementation of the router
func New(controllerSvc controller.API, version string) *Router {
	return &Router{controllerSvc, version}
}

// Healthcheck is just the healthcheck
func (r *Router) Healthcheck(httpw http.ResponseWriter, httpr *http.Request) {
	httpw.WriteHeader(http.StatusOK)
	httpw.Write([]byte(fmt.Sprintf("Here is the version: %s", r.version)))
}

// Login fjdkslfjk
func (r *Router) Login(httpw http.ResponseWriter, httpr *http.Request) {
	var (
		req  LoginRequest
		resp LoginResponse
		err  error
	)
	if err = json.NewDecoder(httpr.Body).Decode(&req); err == nil {
		if err = req.CleanParams(); err == nil {
			if resp.AuthToken, err = r.controllerSvc.Login(req.Email, req.Password); err == nil {
				r.sendSuccess(httpw, resp)
				return
			}
		}
	}
	r.sendFailure(httpw, err, http.StatusInternalServerError)
}

func (r *Router) sendSuccess(httpw http.ResponseWriter, msg interface{}) {
	body, _ := json.Marshal(msg)
	httpw.WriteHeader(http.StatusOK)
	httpw.Write(body)
}

func (r *Router) sendFailure(httpw http.ResponseWriter, err error, code int) {
	type errMsg struct {
		Err error `json:"err"`
	}
	body, _ := json.Marshal(errMsg{Err: err})
	httpw.WriteHeader(code)
	httpw.Write(body)
}
