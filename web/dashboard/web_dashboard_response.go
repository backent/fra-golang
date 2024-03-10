package dashboard

type DashboardResponse struct {
	SummaryAssessment SummaryAssessment `json:"summary_assessement"`
}

type SummaryAssessment struct {
	Release  int `json:"release"`
	Return   int `json:"return"`
	Received int `json:"received"`
	Total    int `json:"total"`
}
