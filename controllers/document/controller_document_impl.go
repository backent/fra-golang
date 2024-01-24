package document

import (
	"context"
	"net/http"
	"strconv"

	"github.com/backent/fra-golang/helpers"
	servicesDocument "github.com/backent/fra-golang/services/document"
	"github.com/backent/fra-golang/web"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/julienschmidt/httprouter"
)

type ControllerDocumentImpl struct {
	servicesDocument.ServiceDocumentInterface
}

func NewControllerDocumentImpl(servicesDocument servicesDocument.ServiceDocumentInterface) ControllerDocumentInterface {
	return &ControllerDocumentImpl{
		ServiceDocumentInterface: servicesDocument,
	}
}

func (implementation *ControllerDocumentImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestCreate
	helpers.DecodeRequest(r, &request)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	createResponse := implementation.ServiceDocumentInterface.Create(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   createResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerDocumentImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestUpdate
	helpers.DecodeRequest(r, &request)

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	createResponse := implementation.ServiceDocumentInterface.Update(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   createResponse,
	}

	helpers.ReturnReponseJSON(w, response)
}
func (implementation *ControllerDocumentImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestDelete

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	implementation.ServiceDocumentInterface.Delete(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerDocumentImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestFindById

	id, err := strconv.Atoi(p.ByName("id"))
	helpers.PanicIfError(err)
	request.Id = id

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	findByIdResponse := implementation.ServiceDocumentInterface.FindById(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   findByIdResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
func (implementation *ControllerDocumentImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestFindAll

	if r.URL.Query().Has("detail") {
		withDetail, err := strconv.ParseBool(r.URL.Query().Get("detail"))
		helpers.PanicIfError(err)
		request.WithDetail = withDetail
	}

	web.SetPagination(&request, r)
	web.SetOrder(&request, r)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	var findAllResponse interface{}
	var total int
	if request.WithDetail {
		findAllResponse, total = implementation.ServiceDocumentInterface.FindAllWithDetail(ctx, request)
	} else {
		findAllResponse, total = implementation.ServiceDocumentInterface.FindAll(ctx, request)
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

func (implementation *ControllerDocumentImpl) GetProductDistinct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestGetProductDistinct

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	documents := implementation.ServiceDocumentInterface.GetProductDistinct(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   documents,
	}

	helpers.ReturnReponseJSON(w, response)

}

func (implementation *ControllerDocumentImpl) Approve(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocument.DocumentRequestApprove
	helpers.DecodeRequest(r, &request)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	implementation.ServiceDocumentInterface.Approve(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)

}
