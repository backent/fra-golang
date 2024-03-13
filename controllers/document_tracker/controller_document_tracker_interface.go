package document_tracker

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerDocumentTrackerInterface interface {
	TrackCount(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
