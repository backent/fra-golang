package risk

import (
	"context"
	"net/http"
	"strconv"

	"github.com/backent/fra-golang/helpers"
	servicesRisk "github.com/backent/fra-golang/services/risk"
	"github.com/backent/fra-golang/web"
	webRisk "github.com/backent/fra-golang/web/risk"
	"github.com/julienschmidt/httprouter"
)

type ControllerRiskImpl struct {
	servicesRisk.ServiceRiskInterface
}

func NewControllerRiskImpl(servicesRisk servicesRisk.ServiceRiskInterface) ControllerRiskInterface {
	return &ControllerRiskImpl{
		ServiceRiskInterface: servicesRisk,
	}
}

func (implementation *ControllerRiskImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webRisk.RiskRequestCreate
	helpers.DecodeRequest(r, &request)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	createResponse := implementation.ServiceRiskInterface.Create(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   createResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerRiskImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webRisk.RiskRequestUpdate
	helpers.DecodeRequest(r, &request)

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	createResponse := implementation.ServiceRiskInterface.Update(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   createResponse,
	}

	helpers.ReturnReponseJSON(w, response)
}
func (implementation *ControllerRiskImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webRisk.RiskRequestDelete

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	implementation.ServiceRiskInterface.Delete(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerRiskImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webRisk.RiskRequestFindById

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	findByIdResponse := implementation.ServiceRiskInterface.FindById(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   findByIdResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerRiskImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webRisk.RiskRequestFindAll

	if r.URL.Query().Has("user") {
		withUser, err := strconv.ParseBool(r.URL.Query().Get("user"))
		helpers.PanicIfError(err)
		request.WithUser = withUser
	}

	web.SetPagination(&request, r)
	web.SetOrder(&request, r)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	var findAllResponse interface{}
	var total int
	if request.WithUser {
		findAllResponse, total = implementation.ServiceRiskInterface.FindAllWithUserDetail(ctx, request)
	} else {
		findAllResponse, total = implementation.ServiceRiskInterface.FindAll(ctx, request)
	}
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
