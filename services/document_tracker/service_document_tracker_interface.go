package document_tracker

import (
	"context"

	"github.com/backent/fra-golang/web/document_tracker"
)

type ServiceDocumentTrackerInterface interface {
	TrackCount(ctx context.Context, request document_tracker.DocumentTrackerRequestTrackCount)
}
