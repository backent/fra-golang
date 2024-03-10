package dashboard

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webDashboard "github.com/backent/fra-golang/web/dashboard"
)

type ServiceDashboardImpl struct {
	DB *sql.DB
	repositoriesUser.RepositoryUserInterface
	repositoriesDocument.RepositoryDocumentInterface
	*middlewares.DashboardMiddleware
}

func NewServiceDashboardImpl(db *sql.DB, repositoriesUser repositoriesUser.RepositoryUserInterface, repositoriesDocument repositoriesDocument.RepositoryDocumentInterface, dashboardMiddleware *middlewares.DashboardMiddleware) ServiceDashboardInterface {
	return &ServiceDashboardImpl{
		DB:                          db,
		RepositoryUserInterface:     repositoriesUser,
		RepositoryDocumentInterface: repositoriesDocument,
		DashboardMiddleware:         dashboardMiddleware,
	}
}

func (implementation *ServiceDashboardImpl) Summary(ctx context.Context, request webDashboard.DashboardRequestSummary) webDashboard.DashboardResponse {

	var dashboardResponse webDashboard.DashboardResponse
	dashboardResponse.SummaryAssessment = getSummaryDocumentAssessment(implementation, ctx, request)

	return dashboardResponse

}

func getSummaryDocumentAssessment(implementation *ServiceDashboardImpl, ctx context.Context, request webDashboard.DashboardRequestSummary) webDashboard.SummaryAssessment {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DashboardMiddleware.Summary(ctx, tx, &request)
	helpers.PanicIfError(err)

	documents, err := implementation.RepositoryDocumentInterface.GetProductCurrentYear(ctx, tx)
	helpers.PanicIfError(err)

	var summary webDashboard.SummaryAssessment
	for _, data := range documents {
		switch data.Action {
		case "approve":
			summary.Release++
		case "reject":
			summary.Return++
		case "submit":
			summary.Received++
		case "update":
			summary.Received++
		}
		summary.Total++
	}

	return summary
}
