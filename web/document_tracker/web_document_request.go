package document_tracker

type DocumentTrackerRequestTrackCount struct {
	Uuid string `json:"uuid" validate:"required"`
	Type string `json:"type" validate:"required,oneof=view search"`
}
