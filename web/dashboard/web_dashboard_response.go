package dashboard

import "github.com/backent/fra-golang/models"

type DashboardResponse struct {
	SummaryAssessment SummaryAssessment                         `json:"summary_assessement"`
	RecentlyViewed    []DashboardResponseTopListDocumentTracker `json:"recently_viewed"`
	TopSearch         []DashboardResponseTopListDocumentTracker `json:"top_search"`
	UserActive        []DashboardResponseUserActive             `json:"user_active"`
}

type SummaryAssessment struct {
	Release  int `json:"release"`
	Return   int `json:"return"`
	Received int `json:"received"`
	Total    int `json:"total"`
}

type DashboardResponseTopListDocumentTracker struct {
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
}

type DashboardResponseUserActive struct {
	Nik  string `json:"nik"`
	Name string `json:"name"`
}

func ModelDocumentTrackerToDashboardResponseTopListDocumentTracker(documentTracker models.DocumentTracker) DashboardResponseTopListDocumentTracker {
	return DashboardResponseTopListDocumentTracker{
		ProductName: documentTracker.DocumentDetail.ProductName,
		Category:    documentTracker.DocumentDetail.Category,
	}
}

func BulkModelDocumentTrackerToDashboardResponseTopListDocumentTracker(documentTrackers []models.DocumentTracker) []DashboardResponseTopListDocumentTracker {
	var bulk []DashboardResponseTopListDocumentTracker

	for _, documentTracker := range documentTrackers {
		bulk = append(bulk, ModelDocumentTrackerToDashboardResponseTopListDocumentTracker(documentTracker))
	}

	return bulk
}

func ModelUserHistoryLoginToDashboardResponseUserActive(userHistoryLogin models.UserHistoryLogin) DashboardResponseUserActive {
	return DashboardResponseUserActive{
		Nik:  userHistoryLogin.UserDetail.Nik,
		Name: userHistoryLogin.UserDetail.Name,
	}
}

func BulkModelUserHistoryLoginToDashboardResponseUserActive(usersHistoryLogin []models.UserHistoryLogin) []DashboardResponseUserActive {
	var bulk []DashboardResponseUserActive

	for _, userHistoryLogin := range usersHistoryLogin {
		bulk = append(bulk, ModelUserHistoryLoginToDashboardResponseUserActive(userHistoryLogin))
	}

	return bulk

}
