package user_registration

import (
	"net/http"

	"github.com/backent/fra-golang/helpers"
	servicesUserRegistration "github.com/backent/fra-golang/services/user_registration"
	"github.com/backent/fra-golang/web"
	webUserRegistration "github.com/backent/fra-golang/web/user_registration"
	"github.com/julienschmidt/httprouter"
)

type ControllerUserRegistrationImpl struct {
	servicesUserRegistration.ServiceUserRegistrationInterface
}

func NewControllerUserRegistrationImpl(servicesUserRegistration servicesUserRegistration.ServiceUserRegistrationInterface) ControllerUserRegistrationInterface {
	return &ControllerUserRegistrationImpl{
		ServiceUserRegistrationInterface: servicesUserRegistration,
	}
}

func (implementation *ControllerUserRegistrationImpl) Apply(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webUserRegistration.UserRegistrationRequestApply
	helpers.DecodeRequest(r, &request)

	ctx := r.Context()

	implementation.ServiceUserRegistrationInterface.Apply(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)

}
