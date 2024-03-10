package dashboard

import (
	"context"

	"github.com/backent/fra-golang/web/dashboard"
)

type ServiceDashboardInterface interface {
	Summary(ctx context.Context, request dashboard.DashboardRequestSummary) dashboard.DashboardResponse
}
