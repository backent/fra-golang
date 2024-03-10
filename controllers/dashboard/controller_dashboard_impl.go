package dashboard

import (
	"context"
	"net/http"

	"github.com/backent/fra-golang/helpers"
	servicesDashboard "github.com/backent/fra-golang/services/dashboard"
	"github.com/backent/fra-golang/web"
	webDashboard "github.com/backent/fra-golang/web/dashboard"
	"github.com/julienschmidt/httprouter"
)

type ControllerDashboardImpl struct {
	servicesDashboard.ServiceDashboardInterface
}

func NewControllerDashboardImpl(servicesDashboard servicesDashboard.ServiceDashboardInterface) ControllerDashboardInterface {
	return &ControllerDashboardImpl{
		ServiceDashboardInterface: servicesDashboard,
	}
}

func (implementation *ControllerDashboardImpl) Summary(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request webDashboard.DashboardRequestSummary

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))

	summary := implementation.ServiceDashboardInterface.Summary(ctx, request)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   summary,
	}

	helpers.ReturnReponseJSON(w, response)

}
