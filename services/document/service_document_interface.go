package document

import (
	"context"

	"github.com/backent/fra-golang/web/document"
)

type ServiceDocumentInterface interface {
	Create(ctx context.Context, request document.DocumentRequestCreate) document.DocumentResponse
	Update(ctx context.Context, request document.DocumentRequestUpdate) document.DocumentResponse
	Delete(ctx context.Context, request document.DocumentRequestDelete)
	FindById(ctx context.Context, request document.DocumentRequestFindById) document.DocumentResponseWithDetail
	FindAll(ctx context.Context, request document.DocumentRequestFindAll) ([]document.DocumentResponse, int)
	FindAllWithDetail(ctx context.Context, request document.DocumentRequestFindAll) ([]document.DocumentResponseWithDetail, int)
	GetProductDistinct(ctx context.Context, request document.DocumentRequestGetProductDistinct) []document.DocumentResponseGetProductDistinct
	Approve(ctx context.Context, request document.DocumentRequestApprove)
	Reject(ctx context.Context, request document.DocumentRequestReject)
	MonitoringList(ctx context.Context, request document.DocumentRequestMonitoringList) ([]document.DocumentResponse, int)
	TrackerProduct(ctx context.Context, request document.DocumentRequestTrackerProduct) []document.DocumentTrackerProduct
}
