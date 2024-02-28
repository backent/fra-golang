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

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

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

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

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

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

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

	web.SetPagination(&request, r)
	web.SetOrder(&request, r)

	if r.URL.Query().Has("status") {
		request.QueryStatus = r.URL.Query().Get("status")
	}

	if r.URL.Query().Has("search") {
		request.QuerySearch = r.URL.Query().Get("search")
	}

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	findAllResponse, total := implementation.ServiceUserInterface.FindAll(ctx, request)

	pagination := web.Pagination{
		Take:  request.GetTake(),
		Skip:  request.GetSkip(),
		Total: total,
	}

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   findAllResponse,
		Extras: pagination,
	}

	helpers.ReturnReponseJSON(w, response)

}

func (implementation *ControllerUserImpl) CurrentUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webUser.UserRequestCurrentUser

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	userResponse := implementation.ServiceUserInterface.CurrentUser(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   userResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
