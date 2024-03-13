package document_tracker

import (
	"context"
	"net/http"

	"github.com/backent/fra-golang/helpers"
	servicesDocumentTracker "github.com/backent/fra-golang/services/document_tracker"
	"github.com/backent/fra-golang/web"
	webDocumentTracker "github.com/backent/fra-golang/web/document_tracker"
	"github.com/julienschmidt/httprouter"
)

type ControllerDocumentTrackerImpl struct {
	servicesDocumentTracker.ServiceDocumentTrackerInterface
}

func NewControllerDocumentTrackerImpl(servicesDocumentTracker servicesDocumentTracker.ServiceDocumentTrackerInterface) ControllerDocumentTrackerInterface {
	return &ControllerDocumentTrackerImpl{
		ServiceDocumentTrackerInterface: servicesDocumentTracker,
	}
}
func (implementation *ControllerDocumentTrackerImpl) TrackCount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDocumentTracker.DocumentTrackerRequestTrackCount
	helpers.DecodeRequest(r, &request)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	implementation.ServiceDocumentTrackerInterface.TrackCount(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)
}
