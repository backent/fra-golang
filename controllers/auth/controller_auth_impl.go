package auth

import (
	"context"
	"net/http"

	"github.com/backent/fra-golang/helpers"
	servicesAuth "github.com/backent/fra-golang/services/auth"
	"github.com/backent/fra-golang/web"
	webAuth "github.com/backent/fra-golang/web/auth"
	"github.com/julienschmidt/httprouter"
)

type ControllerAuthImpl struct {
	servicesAuth.ServiceAuthInterface
}

func NewControllerAuthImpl(servicesAuth servicesAuth.ServiceAuthInterface) ControllerAuthInterface {
	return &ControllerAuthImpl{
		ServiceAuthInterface: servicesAuth,
	}
}

func (implementation *ControllerAuthImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var loginRequest webAuth.LoginRequest
	helpers.DecodeRequest(r, &loginRequest)

	ctx := context.Background()
	response := implementation.ServiceAuthInterface.Login(ctx, loginRequest)

	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   response,
	}

	helpers.ReturnReponseJSON(w, webResponse)
}
func (implementation *ControllerAuthImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
