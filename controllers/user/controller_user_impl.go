package user

import (
	"context"
	"net/http"
	"strconv"

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
	var request webUser.UserRequestUpdate
	helpers.DecodeRequest(r, &request)

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := r.Context()

	createResponse := implementation.ServiceUserInterface.Update(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   createResponse,
	}

	helpers.ReturnReponseJSON(w, response)
}
func (implementation *ControllerUserImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webUser.UserRequestDelete

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := r.Context()

	implementation.ServiceUserInterface.Delete(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerUserImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webUser.UserRequestFindById

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := r.Context()

	findByIdResponse := implementation.ServiceUserInterface.FindById(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   findByIdResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerUserImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webUser.UserRequestFindAll

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	findAllResponse := implementation.ServiceUserInterface.FindAll(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   findAllResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
