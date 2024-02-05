package notification

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerNotificationInterface interface {
	ReadAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
