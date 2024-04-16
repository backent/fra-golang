package document

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerDocumentInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetProductDistinct(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Approve(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Reject(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	MonitoringList(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	TrackerProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	SearchGlobal(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UploadFinal(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	ServeFile(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
