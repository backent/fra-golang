package risk

import (
	"context"

	"github.com/backent/fra-golang/web/risk"
)

type ServiceRiskInterface interface {
	Create(ctx context.Context, request risk.RiskRequestCreate) risk.RiskResponse
	Update(ctx context.Context, request risk.RiskRequestUpdate) risk.RiskResponse
	Delete(ctx context.Context, request risk.RiskRequestDelete)
	FindById(ctx context.Context, request risk.RiskRequestFindById) risk.RiskResponse
	FindAll(ctx context.Context, request risk.RiskRequestFindAll) ([]risk.RiskResponse, int)
	FindAllWithUserDetail(ctx context.Context, request risk.RiskRequestFindAll) ([]risk.RiskResponseWithUserDetail, int)
}
