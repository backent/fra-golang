package document

import (
	"time"

	"github.com/backent/fra-golang/models"
)

type DocumentResponse struct {
	Id          int       `json:"id"`           // id
	Uuid        string    `json:"uuid"`         // uuid
	CreatedBy   int       `json:"created_by"`   // created_by
	ActionBy    int       `json:"action_by"`    // action_by
	Action      string    `json:"action"`       // action
	ProductName string    `json:"product_name"` // product_name
	CreatedAt   time.Time `json:"created_at"`   // created_at
	UpdatedAt   time.Time `json:"updated_at"`   // updated_at
}

func DocumentModelToDocumentResponse(document models.Document) DocumentResponse {
	return DocumentResponse{
		Id:          document.Id,
		Uuid:        document.Uuid,
		CreatedBy:   document.CreatedBy,
		ActionBy:    document.ActionBy,
		Action:      document.Action,
		ProductName: document.ProductName,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}
}

func BulkDocumentModelToDocumentResponse(documents []models.Document) []DocumentResponse {
	var bulk []DocumentResponse
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponse(document))
	}
	return bulk
}

type DocumentResponseWithDetail struct {
	Id          int       `json:"id"`           // id
	Uuid        string    `json:"uuid"`         // uuid
	CreatedBy   int       `json:"created_by"`   // created_by
	ActionBy    int       `json:"action_by"`    // action_by
	Action      string    `json:"action"`       // action
	ProductName string    `json:"product_name"` // product_name
	CreatedAt   time.Time `json:"created_at"`   // created_at
	UpdatedAt   time.Time `json:"updated_at"`   // updated_at

	UserDetail            userResponse              `json:"user_detail"`             // user detail
	RiskDetail            []riskResponse            `json:"risk_detail"`             // risk detail
	RelatedDocumentDetail []RelatedDocumentResponse `json:"related_document_detail"` // related document detail
}

type userResponse struct {
	Id   int    `json:"id"`
	Nik  string `json:"nik"`
	Name string `json:"name"`
}

type riskResponse struct {
	Id                     int       `json:"id"`                       // id
	DocumentId             int       `json:"document_id"`              // document_id
	RiskName               string    `json:"risk_name"`                // risk_name
	FraudSchema            string    `json:"fraud_schema"`             // fraud_schema
	FraudMotive            string    `json:"fraud_motive"`             // fraud_motive
	FraudTechnique         string    `json:"fraud_technique"`          // fraud_technique
	RiskSource             string    `json:"risk_source"`              // risk_source
	RootCause              string    `json:"root_cause"`               // root_cause
	BisproControlProcedure string    `json:"bispro_control_procedure"` // bispro_control_procedure
	QualitativeImpact      string    `json:"qualitative_impact"`       // qualitative_impact
	LikehoodJustification  string    `json:"likehood_justification"`   // likehood_justification
	ImpactJustification    string    `json:"impact_justification"`     // impact_justification
	StartegyAgreement      string    `json:"strategy_agreement"`       // strategy_agreement
	StrategyRecomendation  string    `json:"strategy_recomendation"`   // strategy_recomendation
	AssessmentLikehood     string    `json:"assessment_likehood"`      // assessment_likehood
	AssessmentImpact       string    `json:"assessment_impact"`        // assessment_impact
	AssessmentRiskLevel    string    `json:"assessment_risk_level"`    // assessment_risk_level
	CreatedAt              time.Time `json:"created_at"`               // created_at
	UpdatedAt              time.Time `json:"updated_at"`               // updated_at

	RejectNoteDetail *RejectNoteResponse `json:"reject_note_detail"`
}

type RejectNoteResponse struct {
	Id                     int    `json:"id"`                       // id
	DocumentId             int    `json:"document_id"`              // document_id
	RiskId                 int    `json:"risk_id"`                  // risk_id
	Fraud                  string `json:"fraud"`                    // fraud
	RiskSource             string `json:"risk_source"`              // risk_source
	RootCause              string `json:"root_cause"`               // root_cause
	BisproControlProcedure string `json:"bispro_control_procedure"` // bispro_control_procedure
	QualitativeImpact      string `json:"qualitative_impact"`       // qualitative_impact
	Assessment             string `json:"assessment"`               // assessment
	Justification          string `json:"justification"`            // justification
	Strategy               string `json:"strategy"`                 // strategy
}

type RelatedDocumentResponse struct {
	Id        int       `json:"id"`
	CreatedBy time.Time `json:"created_by"`
}

func DocumentModelToDocumentResponseWithDetail(document models.Document) DocumentResponseWithDetail {
	var riskDetail []riskResponse
	if temp := riskBulkToRiskResponseBulk(document.RiskDetail); len(temp) > 0 {
		riskDetail = temp
	} else {
		riskDetail = make([]riskResponse, 0)
	}

	var relatedDocumentDetail []RelatedDocumentResponse
	if len(document.RelatedDocumentDetail) > 0 {
		relatedDocumentDetail = bulkRelatedDocumentToRelatedDocumentResponse(document.RelatedDocumentDetail)
	} else {
		relatedDocumentDetail = make([]RelatedDocumentResponse, 0)
	}
	return DocumentResponseWithDetail{
		Id:                    document.Id,
		Uuid:                  document.Uuid,
		CreatedBy:             document.CreatedBy,
		ActionBy:              document.ActionBy,
		Action:                document.Action,
		ProductName:           document.ProductName,
		CreatedAt:             document.CreatedAt,
		UpdatedAt:             document.UpdatedAt,
		UserDetail:            userModelToUserResponse(document.UserDetail),
		RiskDetail:            riskDetail,
		RelatedDocumentDetail: relatedDocumentDetail,
	}
}

func BulkDocumentModelToDocumentResponseWithDetail(documents []models.Document) []DocumentResponseWithDetail {
	var bulk []DocumentResponseWithDetail
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponseWithDetail(document))
	}
	return bulk
}

func userModelToUserResponse(user models.User) userResponse {
	return userResponse{
		Id:   user.Id,
		Name: user.Name,
		Nik:  user.Nik,
	}
}

func riskToRiskResponse(risk models.Risk) riskResponse {
	var rejectNoteResponse *RejectNoteResponse
	if risk.RejectNoteDetail.Id != 0 {
		rejectNoteResponse = rejectNoteToRejectNoteResponse(risk.RejectNoteDetail)
	} else {
		rejectNoteResponse = nil
	}

	return riskResponse{
		Id:                     risk.Id,
		DocumentId:             risk.DocumentId,
		RiskName:               risk.RiskName,
		FraudSchema:            risk.FraudSchema,
		FraudMotive:            risk.FraudMotive,
		FraudTechnique:         risk.FraudTechnique,
		RiskSource:             risk.RiskSource,
		RootCause:              risk.RootCause,
		BisproControlProcedure: risk.BisproControlProcedure,
		QualitativeImpact:      risk.QualitativeImpact,
		LikehoodJustification:  risk.LikehoodJustification,
		ImpactJustification:    risk.ImpactJustification,
		StartegyAgreement:      risk.StartegyAgreement,
		StrategyRecomendation:  risk.StrategyRecomendation,
		AssessmentLikehood:     risk.AssessmentLikehood,
		AssessmentImpact:       risk.AssessmentImpact,
		AssessmentRiskLevel:    risk.AssessmentRiskLevel,
		CreatedAt:              risk.CreatedAt,
		UpdatedAt:              risk.UpdatedAt,
		RejectNoteDetail:       rejectNoteResponse,
	}
}

func riskBulkToRiskResponseBulk(risks []models.Risk) []riskResponse {
	var risksResponse []riskResponse
	for _, risk := range risks {
		risksResponse = append(risksResponse, riskToRiskResponse(risk))
	}
	return risksResponse
}

func relatedDocumentToRelatedDocumentResponse(relatedDocument models.RelatedDocument) RelatedDocumentResponse {
	return RelatedDocumentResponse{
		Id:        relatedDocument.Id,
		CreatedBy: relatedDocument.CreatedAt,
	}
}

func bulkRelatedDocumentToRelatedDocumentResponse(relatedDocument []models.RelatedDocument) []RelatedDocumentResponse {
	var bulk []RelatedDocumentResponse
	for _, item := range relatedDocument {
		bulk = append(bulk, relatedDocumentToRelatedDocumentResponse(item))
	}

	return bulk
}

func rejectNoteToRejectNoteResponse(rejectNote models.RejectNote) *RejectNoteResponse {
	return &RejectNoteResponse{
		Id:                     rejectNote.Id,
		DocumentId:             rejectNote.DocumentId,
		RiskId:                 rejectNote.RiskId,
		Fraud:                  rejectNote.Fraud,
		RiskSource:             rejectNote.RiskSource,
		RootCause:              rejectNote.RootCause,
		BisproControlProcedure: rejectNote.BisproControlProcedure,
		QualitativeImpact:      rejectNote.QualitativeImpact,
		Assessment:             rejectNote.Assessment,
		Justification:          rejectNote.Justification,
		Strategy:               rejectNote.Strategy,
	}
}

type DocumentResponseGetProductDistinct struct {
	Id          int    `json:"id"`
	ProductName string `json:"product_name"`
}

func DocumentModelToDocumentResponseGetProductDistinct(document models.Document) DocumentResponseGetProductDistinct {
	return DocumentResponseGetProductDistinct{
		Id:          document.Id,
		ProductName: document.ProductName,
	}
}

func BulkDocumentModelToBulkDocumentResponseGetProductDistinct(documents []models.Document) []DocumentResponseGetProductDistinct {
	var bulkDocumentResponseGetProductDistinct []DocumentResponseGetProductDistinct
	for _, document := range documents {
		bulkDocumentResponseGetProductDistinct = append(bulkDocumentResponseGetProductDistinct, DocumentModelToDocumentResponseGetProductDistinct(document))
	}
	return bulkDocumentResponseGetProductDistinct
}
