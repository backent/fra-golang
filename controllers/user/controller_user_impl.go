package user

import (
	"net/http"

	"github.com/backent/fra-golang/helpers"
	servicesUser "github.com/backent/fra-golang/services/user"
	"github.com/backent/fra-golang/web"
	webUser "github.com/backent/fra-golang/web/user"
	"github.com/julienschmidt/httprouter"
)

type ControllerUserImpl struct {
	servicesUser.ServiceUserInterface
}

func NewControllerUserImpl(servicesUser servicesUser.ServiceUserInterface) ControllerUserInterface {
	return &ControllerUserImpl{
		ServiceUserInterface: servicesUser,
	}
}

func (implementation *ControllerUserImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webUser.UserRequestCreate
	helpers.DecodeRequest(r, &request)

	ctx := r.Context()

	createResponse := implementation.ServiceUserInterface.Create(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   createResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerUserImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
