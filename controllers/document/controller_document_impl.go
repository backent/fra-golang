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

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	findAllResponse := implementation.ServiceDocumentInterface.FindAll(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   findAllResponse,
	}

	helpers.ReturnReponseJSON(w, response)

}
