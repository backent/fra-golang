package notification

import (
	"context"
	"net/http"

	"github.com/backent/fra-golang/helpers"
	servicesNotification "github.com/backent/fra-golang/services/notification"
	"github.com/backent/fra-golang/web"
	webNotification "github.com/backent/fra-golang/web/notification"
	"github.com/julienschmidt/httprouter"
)

type ControllerNotificationImpl struct {
	servicesNotification.ServiceNotificationInterface
}

func NewControllerNotificationImpl(servicesNotification servicesNotification.ServiceNotificationInterface) ControllerNotificationInterface {
	return &ControllerNotificationImpl{
		ServiceNotificationInterface: servicesNotification,
	}
}
func (implementation *ControllerNotificationImpl) ReadAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webNotification.NotificationRequestReadAll

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	implementation.ServiceNotificationInterface.ReadAll(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnReponseJSON(w, response)
}
func (implementation *ControllerNotificationImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webNotification.NotificationRequestFindAll

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	notifications := implementation.ServiceNotificationInterface.FindAll(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   notifications,
	}

	helpers.ReturnReponseJSON(w, response)
}
