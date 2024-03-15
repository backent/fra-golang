package dashboard

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesDocumentTracker "github.com/backent/fra-golang/repositories/document_tracker"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webDashboard "github.com/backent/fra-golang/web/dashboard"
)

type ServiceDashboardImpl struct {
	DB *sql.DB
	repositoriesUser.RepositoryUserInterface
	repositoriesDocument.RepositoryDocumentInterface
	repositoriesDocumentTracker.RepositoryDocumentTrackerInterface
	*middlewares.DashboardMiddleware
}

func NewServiceDashboardImpl(db *sql.DB, repositoriesUser repositoriesUser.RepositoryUserInterface, repositoriesDocument repositoriesDocument.RepositoryDocumentInterface, repositoriesDocumentTracker repositoriesDocumentTracker.RepositoryDocumentTrackerInterface, dashboardMiddleware *middlewares.DashboardMiddleware) ServiceDashboardInterface {
	return &ServiceDashboardImpl{
		DB:                                 db,
		RepositoryUserInterface:            repositoriesUser,
		RepositoryDocumentInterface:        repositoriesDocument,
		RepositoryDocumentTrackerInterface: repositoriesDocumentTracker,
		DashboardMiddleware:                dashboardMiddleware,
	}
}

func (implementation *ServiceDashboardImpl) Summary(ctx context.Context, request webDashboard.DashboardRequestSummary) webDashboard.DashboardResponse {

	var dashboardResponse webDashboard.DashboardResponse

	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DashboardMiddleware.Summary(ctx, tx, &request)

	chanSummaryAssessmentErr := make(chan error)
	chanMostViewedErr := make(chan error)
	chanTopSearchErr := make(chan error)

	go func() {
		dashboardResponse.SummaryAssessment, err = getSummaryDocumentAssessment(implementation, ctx, request)
		chanSummaryAssessmentErr <- err
	}()

	go func() {
		dashboardResponse.RecentlyViewed, err = getRecentlyMostViewed(implementation, ctx, request)
		chanMostViewedErr <- err

	}()

	go func() {
		dashboardResponse.TopSearch, err = getTopSearched(implementation, ctx, request)
		chanTopSearchErr <- err
	}()

	summaryAssessmentErr := <-chanSummaryAssessmentErr
	mostViewedErr := <-chanMostViewedErr
	topSearchErr := <-chanTopSearchErr

	helpers.PanicIfError(summaryAssessmentErr)
	helpers.PanicIfError(mostViewedErr)
	helpers.PanicIfError(topSearchErr)

	return dashboardResponse

}

func getSummaryDocumentAssessment(implementation *ServiceDashboardImpl, ctx context.Context, request webDashboard.DashboardRequestSummary) (webDashboard.SummaryAssessment, error) {

	var summary webDashboard.SummaryAssessment

	tx, err := implementation.DB.Begin()
	if err != nil {
		return summary, err
	}
	documents, err := implementation.RepositoryDocumentInterface.GetProductCurrentYear(ctx, tx)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return summary, errRollback
		}
		return summary, errRollback
	}

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

	errCommit := tx.Commit()
	if errCommit != nil {
		return summary, err
	}

	return summary, nil
}

func getRecentlyMostViewed(implementation *ServiceDashboardImpl, ctx context.Context, request webDashboard.DashboardRequestSummary) ([]webDashboard.DashboardResponseTopListDocumentTracker, error) {

	tx, err := implementation.DB.Begin()
	if err != nil {
		return nil, err
	}

	currentYear := strconv.Itoa(time.Time.Year(time.Now()))
	documentTrackers, _, err := implementation.RepositoryDocumentTrackerInterface.FindAll(ctx, tx, 5, 0, "viewed_count", "DESC", currentYear)
	if err != nil {
		return nil, err
	}

	if len(documentTrackers) > 0 {
		return webDashboard.BulkModelDocumentTrackerToDashboardResponseTopListDocumentTracker(documentTrackers), nil
	}

	return []webDashboard.DashboardResponseTopListDocumentTracker{}, nil

}

func getTopSearched(implementation *ServiceDashboardImpl, ctx context.Context, request webDashboard.DashboardRequestSummary) ([]webDashboard.DashboardResponseTopListDocumentTracker, error) {

	tx, err := implementation.DB.Begin()
	if err != nil {
		return nil, err
	}

	currentYear := strconv.Itoa(time.Time.Year(time.Now()))
	documentTrackers, _, err := implementation.RepositoryDocumentTrackerInterface.FindAll(ctx, tx, 5, 0, "searched_count", "DESC", currentYear)
	if err != nil {
		return nil, err
	}

	if len(documentTrackers) > 0 {
		return webDashboard.BulkModelDocumentTrackerToDashboardResponseTopListDocumentTracker(documentTrackers), nil
	}

	return []webDashboard.DashboardResponseTopListDocumentTracker{}, nil

}
