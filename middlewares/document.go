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

	// make sure the action is lowercase
	request.Action = strings.ToLower(request.Action)

	ValidateUserPermission(ctx, tx, implementation.RepositoryUserInterface, userId, request.Action)

	for index := range request.Risks {
		request.Risks[index].AssessmentImpact = strings.ToLower(request.Risks[index].AssessmentImpact)
		request.Risks[index].AssessmentLikehood = strings.ToLower(request.Risks[index].AssessmentLikehood)
		request.Risks[index].AssessmentRiskLevel = strings.ToLower(request.Risks[index].AssessmentRiskLevel)
	}

	if request.Uuid != "" {
		documents, err := implementation.RepositoryDocumentInterface.FindByUUID(ctx, tx, request.Uuid)

		// check if uuid is present, is there any document with that uuid or not
		if err != nil || len(documents) == 0 {
			panic(exceptions.NewNotFoundError(err.Error()))
		}

		// check if uuid is present, is the latest document older / newer than now
		if documents[0].Id != request.Id {
			panic(exceptions.NewConflictError("version mismatch"))
		}

		if strings.ToLower(request.Action) == "submit" {
			nonDraftDocument, err := implementation.RepositoryDocumentInterface.GetNonDraftProductByUUID(ctx, tx, request.Uuid)
			helpers.PanicIfError(err)
			if len(nonDraftDocument) > 0 {
				request.Action = "update"
			}
		}
		request.CreatedBy = documents[0].CreatedBy
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

	switch user.Role {
	case "guest":
		request.QueryAction = "approve,final"
	case "reviewer":
		request.QueryAction = "submit,approve,final,reject,update,draft"
	case "superadmin":
		request.QueryAction = "submit,approve,final,reject,update,draft"
	case "author":
		if user.Unit == request.QueryCategory {
			request.QueryAction = "submit,approve,final,reject,update,draft"
		} else {
			request.QueryAction = "approve,final"
		}
	default:
		// just to make user with no role cannot see anything
		// because there is no document with action "none"
		request.QueryAction = "none"
	}

}

func (implementation *DocumentMiddleware) GetProductDistinct(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestGetProductDistinct) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *DocumentMiddleware) Approve(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestApprove) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	ValidateUserPermission(ctx, tx, implementation.RepositoryUserInterface, userId, "approve")

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	documents, err := implementation.RepositoryDocumentInterface.FindByUUID(ctx, tx, document.Uuid)
	helpers.PanicIfError(err)
	if documents[0].Id != request.Id {
		panic(exceptions.NewConflictError("version mismatch"))
	}

	document.Action = "approve"
	document.ActionBy = userId
	request.Document = document

}

func (implementation *DocumentMiddleware) Reject(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestReject) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	ValidateUserPermission(ctx, tx, implementation.RepositoryUserInterface, userId, "reject")

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	documents, err := implementation.RepositoryDocumentInterface.FindByUUID(ctx, tx, document.Uuid)
	helpers.PanicIfError(err)
	if documents[0].Id != request.Id {
		panic(exceptions.NewConflictError("version mismatch"))
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
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)

	user, _ := implementation.RepositoryUserInterface.FindById(ctx, tx, userId)

	switch user.Role {
	case "guest":
		request.QueryAction = "submit,approve,reject,update,final"
	default:
		request.QueryAction = "approve,final"
	}

}

func (implementation *DocumentMiddleware) TrackerProduct(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestTrackerProduct) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *DocumentMiddleware) SearchGlobal(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestSearchGlobal) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *DocumentMiddleware) UploadFinal(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestUploadFinal) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	ValidateUserPermission(ctx, tx, implementation.RepositoryUserInterface, userId, "approve")

	if request.Id == 0 {
		panic(exceptions.NewBadRequestError("id is required"))
	}

	if request.File == nil {
		panic(exceptions.NewBadRequestError("file is required"))
	}

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	documents, err := implementation.RepositoryDocumentInterface.FindByUUID(ctx, tx, document.Uuid)
	helpers.PanicIfError(err)
	if documents[0].Id != request.Id {
		panic(exceptions.NewConflictError("version mismatch"))
	}

	if document.Action != "approve" {
		panic(exceptions.NewBadRequestError("only approved document can save as final"))
	}

	document.Action = "final"
	document.ActionBy = userId
	request.Document = document

}
