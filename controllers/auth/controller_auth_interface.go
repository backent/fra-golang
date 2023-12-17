package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerAuthInterface interface {
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Register(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
