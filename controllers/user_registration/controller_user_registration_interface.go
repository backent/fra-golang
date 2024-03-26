package user_registration

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerUserRegistrationInterface interface {
	Apply(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Approve(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Reject(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	CheckUserLDAP(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
