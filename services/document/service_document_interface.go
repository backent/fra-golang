package document

import (
	"context"

	"github.com/backent/fra-golang/web/document"
)

type ServiceDocumentInterface interface {
	Create(ctx context.Context, request document.DocumentRequestCreate) document.DocumentResponse
	Update(ctx context.Context, request document.DocumentRequestUpdate) document.DocumentResponse
	Delete(ctx context.Context, request document.DocumentRequestDelete)
	FindById(ctx context.Context, request document.DocumentRequestFindById) document.DocumentResponse
	FindAll(ctx context.Context, request document.DocumentRequestFindAll) ([]document.DocumentResponse, int)
	FindAllWithDetail(ctx context.Context, request document.DocumentRequestFindAll) ([]document.DocumentResponseWithDetail, int)
}
