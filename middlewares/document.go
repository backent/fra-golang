package middlewares

import (
	"context"
	"database/sql"
	"strings"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DocumentMiddleware struct {
	Validate *validator.Validate
	repositoriesDocument.RepositoryDocumentInterface
	repositoriesAuth.RepositoryAuthInterface
	repositoriesUser.RepositoryUserInterface
	repositoriesRisk.RepositoryRiskInterface
}

func NewDocumentMiddleware(
	validator *validator.Validate,
	repositoriesDocument repositoriesDocument.RepositoryDocumentInterface,
	repositoriesAuth repositoriesAuth.RepositoryAuthInterface,
	repositoriesUser repositoriesUser.RepositoryUserInterface,
	repositoriesRisk repositoriesRisk.RepositoryRiskInterface,
) *DocumentMiddleware {
	return &DocumentMiddleware{
		Validate:                    validator,
		RepositoryDocumentInterface: repositoriesDocument,
		RepositoryAuthInterface:     repositoriesAuth,
		RepositoryUserInterface:     repositoriesUser,
		RepositoryRiskInterface:     repositoriesRisk,
	}
}

func (implementation *DocumentMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestCreate) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	for index := range request.Risks {
		request.Risks[index].AssessmentImpact = strings.ToLower(request.Risks[index].AssessmentImpact)
		request.Risks[index].AssessmentLikehood = strings.ToLower(request.Risks[index].AssessmentLikehood)
		request.Risks[index].AssessmentRiskLevel = strings.ToLower(request.Risks[index].AssessmentRiskLevel)
	}

	if request.Uuid != "" {
		document, err := implementation.RepositoryDocumentInterface.FindByUUID(ctx, tx, request.Uuid)
		if err != nil {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		request.CreatedBy = document.CreatedBy
	} else {
		request.Uuid = uuid.New().String()
		request.CreatedBy = userId
	}

	request.ActionBy = userId
}

func (implementation *DocumentMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestUpdate) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.Id = document.Id
	request.Uuid = document.Uuid
	request.CreatedBy = document.CreatedBy
	request.ActionBy = userId
}

func (implementation *DocumentMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestDelete) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *DocumentMiddleware) FindById(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestFindById) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *DocumentMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestFindAll) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	user, _ := implementation.RepositoryUserInterface.FindById(ctx, tx, userId)
	if user.Role == "author" {
		request.CreatedBy = user.Id
	} else {
		request.QueryAction = "submit,approve,reject,update"
	}

}

func (implementation *DocumentMiddleware) GetProductDistinct(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestGetProductDistinct) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *DocumentMiddleware) Approve(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestApprove) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	document.Action = "approve"
	document.ActionBy = userId
	request.Document = document

}

func (implementation *DocumentMiddleware) Reject(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestReject) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	for idx := range request.RejectNote {
		risk, err := implementation.RepositoryRiskInterface.FindById(ctx, tx, request.RejectNote[idx].RiskId)
		if err != nil {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		if document.Id != risk.DocumentId {
			panic(exceptions.NewBadRequestError("mismatch document id and risk id"))
		}
	}

	request.Document = document
	request.Document.Action = "reject"
	request.Document.ActionBy = userId

}

func (implementation *DocumentMiddleware) MonitoringList(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestMonitoringList) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	request.QueryAction = "submit,approve,reject,update"
}
