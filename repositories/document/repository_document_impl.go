package document

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/backent/fra-golang/models"
)

type RepositoryDocumentImpl struct {
}

func NewRepositoryDocumentImpl() RepositoryDocumentInterface {
	return &RepositoryDocumentImpl{}
}

func (implementation *RepositoryDocumentImpl) Create(ctx context.Context, tx *sql.Tx, document models.Document) (models.Document, error) {
	query := fmt.Sprintf(`INSERT INTO %s (
		document_id,
		user_id,
		risk_name,
		fraud_schema,
		fraud_motive,
		fraud_technique,
		risk_source,
		root_cause,
		bispro_control_procedure,
		qualitative_impact,
		likehood_justification,
		impact_justification,
		strategy_agreement,
		strategy_recomendation,
		assessment_likehood,
		assessment_impact,
		assessment_risk_level,
		action,
		action_by
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `, models.DocumentTable)
	result, err := tx.ExecContext(ctx, query,
		document.DocumentId,
		document.UserId,
		document.RiskName,
		document.FraudSchema,
		document.FraudMotive,
		document.FraudTechnique,
		document.RiskSource,
		document.RootCause,
		document.BisproControlProcedure,
		document.QualitativeImpact,
		document.LikehoodJustification,
		document.ImpactJustification,
		document.StartegyAgreement,
		document.StrategyRecomendation,
		document.AssessmentLikehood,
		document.AssessmentImpact,
		document.AssessmentRiskLevel,
		document.Action,
		document.ActionBy,
	)
	if err != nil {
		return document, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return document, err
	}

	document.Id = int(id)

	return document, nil

}
func (implementation *RepositoryDocumentImpl) Update(ctx context.Context, tx *sql.Tx, document models.Document) (models.Document, error) {
	query := fmt.Sprintf(`UPDATE  %s SET 
		risk_name = ?,
		fraud_schema = ?,
		fraud_motive = ?,
		fraud_technique = ?,
		risk_source = ?,
		root_cause = ?,
		bispro_control_procedure = ?,
		qualitative_impact = ?,
		likehood_justification = ?,
		impact_justification = ?,
		strategy_agreement = ?,
		strategy_recomendation = ?,
		assessment_likehood = ?,
		assessment_impact = ?,
		assessment_risk_level = ?,
		action = ?,
		action_by = ?
	WHERE id = ?`, models.DocumentTable)
	_, err := tx.ExecContext(ctx, query,

		document.RiskName,
		document.FraudSchema,
		document.FraudMotive,
		document.FraudTechnique,
		document.RiskSource,
		document.RootCause,
		document.BisproControlProcedure,
		document.QualitativeImpact,
		document.LikehoodJustification,
		document.ImpactJustification,
		document.StartegyAgreement,
		document.StrategyRecomendation,
		document.AssessmentLikehood,
		document.AssessmentImpact,
		document.AssessmentRiskLevel,
		document.Action,
		document.ActionBy,

		document.Id)
	if err != nil {
		return document, err
	}

	return document, nil
}
func (implementation *RepositoryDocumentImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := fmt.Sprintf("DELETE FROM  %s  WHERE id = ?", models.DocumentTable)
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (implementation *RepositoryDocumentImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.Document, error) {
	var document models.Document

	query := fmt.Sprintf(`SELECT 
		id,
		document_id,
		user_id,
		risk_name,
		fraud_schema,
		fraud_motive,
		fraud_technique,
		risk_source,
		root_cause,
		bispro_control_procedure,
		qualitative_impact,
		likehood_justification,
		impact_justification,
		strategy_agreement,
		strategy_recomendation,
		assessment_likehood,
		assessment_impact,
		assessment_risk_level,
		action,
		action_by,
		created_at,
		updated_at
	FROM %s WHERE id = ?`, models.DocumentTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return document, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&document.Id,
			&document.DocumentId,
			&document.UserId,
			&document.RiskName,
			&document.FraudSchema,
			&document.FraudMotive,
			&document.FraudTechnique,
			&document.RiskSource,
			&document.RootCause,
			&document.BisproControlProcedure,
			&document.QualitativeImpact,
			&document.LikehoodJustification,
			&document.ImpactJustification,
			&document.StartegyAgreement,
			&document.StrategyRecomendation,
			&document.AssessmentLikehood,
			&document.AssessmentImpact,
			&document.AssessmentRiskLevel,
			&document.Action,
			&document.ActionBy,
			&document.CreatedAt,
			&document.UpdatedAt,
		)
		if err != nil {
			return document, err
		}
	} else {
		return document, errors.New("not found document")
	}

	return document, nil
}

func (implementation *RepositoryDocumentImpl) FindByDocumentId(ctx context.Context, tx *sql.Tx, documentId string) (models.Document, error) {
	var document models.Document
	query := fmt.Sprintf(`SELECT 
		id,
		document_id,
		user_id,
		risk_name,
		fraud_schema,
		fraud_motive,
		fraud_technique,
		risk_source,
		root_cause,
		bispro_control_procedure,
		qualitative_impact,
		likehood_justification,
		impact_justification,
		strategy_agreement,
		strategy_recomendation,
		assessment_likehood,
		assessment_impact,
		assessment_risk_level,
		action,
		action_by,
		created_at,
		updated_at
	FROM %s WHERE document_id = ?`, models.DocumentTable)
	rows, err := tx.QueryContext(ctx, query, documentId)
	if err != nil {
		return document, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&document.Id,
			&document.DocumentId,
			&document.UserId,
			&document.RiskName,
			&document.FraudSchema,
			&document.FraudMotive,
			&document.FraudTechnique,
			&document.RiskSource,
			&document.RootCause,
			&document.BisproControlProcedure,
			&document.QualitativeImpact,
			&document.LikehoodJustification,
			&document.ImpactJustification,
			&document.StartegyAgreement,
			&document.StrategyRecomendation,
			&document.AssessmentLikehood,
			&document.AssessmentImpact,
			&document.AssessmentRiskLevel,
			&document.Action,
			&document.ActionBy,
			&document.CreatedAt,
			&document.UpdatedAt,
		)
		if err != nil {
			return document, err
		}
	} else {
		return document, errors.New("not found document")
	}

	return document, nil
}

func (implementation *RepositoryDocumentImpl) FindAllWithUserDetail(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Document, int, error) {
	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s
		)
		SELECT 
		a.id,
		a.document_id,
		a.user_id,
		a.risk_name,
		a.fraud_schema,
		a.fraud_motive,
		a.fraud_technique,
		a.risk_source,
		a.root_cause,
		a.bispro_control_procedure,
		a.qualitative_impact,
		a.likehood_justification,
		a.impact_justification,
		a.strategy_agreement,
		a.strategy_recomendation,
		a.assessment_likehood,
		a.assessment_impact,
		a.assessment_risk_level,
		a.action,
		a.action_by,
		a.created_at,
		a.updated_at,
		b.id,
		b.name,
		b.nik,
		c.count
	FROM (SELECT * FROM main_table ORDER BY %s %s  LIMIT ?, ?) a LEFT JOIN %s b ON a.user_id = b.id LEFT JOIN (SELECT COUNT(*) as count FROM main_table) c ON true `, models.DocumentTable, orderBy, orderDirection, models.UserTable)
	rows, err := tx.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var totalDocument int
	var documents []*models.Document
	documentsMap := make(map[int]*models.Document)

	for rows.Next() {
		var document models.Document
		var user models.User

		err = rows.Scan(&document.Id,
			&document.DocumentId,
			&document.UserId,
			&document.RiskName,
			&document.FraudSchema,
			&document.FraudMotive,
			&document.FraudTechnique,
			&document.RiskSource,
			&document.RootCause,
			&document.BisproControlProcedure,
			&document.QualitativeImpact,
			&document.LikehoodJustification,
			&document.ImpactJustification,
			&document.StartegyAgreement,
			&document.StrategyRecomendation,
			&document.AssessmentLikehood,
			&document.AssessmentImpact,
			&document.AssessmentRiskLevel,
			&document.Action,
			&document.ActionBy,
			&document.CreatedAt,
			&document.UpdatedAt,
			&user.Id,
			&user.Name,
			&user.Nik,
			&totalDocument,
		)
		if err != nil {
			return nil, 0, err
		}

		item, found := documentsMap[document.Id]
		if !found {
			item = &document
			documentsMap[document.Id] = item
			documents = append(documents, item)
		}
		item.UserDetail = user
	}

	var documentsReturn []models.Document
	for _, document := range documents {
		documentsReturn = append(documentsReturn, *document)
	}

	return documentsReturn, totalDocument, nil
}

func (implementation *RepositoryDocumentImpl) FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Document, int, error) {
	var documents []models.Document

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s
		)
		SELECT 
		a.id,
		a.document_id,
		a.user_id,
		a.risk_name,
		a.fraud_schema,
		a.fraud_motive,
		a.fraud_technique,
		a.risk_source,
		a.root_cause,
		a.bispro_control_procedure,
		a.qualitative_impact,
		a.likehood_justification,
		a.impact_justification,
		a.strategy_agreement,
		a.strategy_recomendation,
		a.assessment_likehood,
		a.assessment_impact,
		a.assessment_risk_level,
		a.action,
		a.action_by,
		a.created_at,
		a.updated_at,
		b.count
	FROM (SELECT * FROM main_table ORDER BY %s %s LIMIT ?, ?) a LEFT JOIN (SELECT COUNT(*) as count FROM main_table) b ON true`, models.DocumentTable, orderBy, orderDirection)
	rows, err := tx.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var totalDocument int
	for rows.Next() {
		var document models.Document
		err = rows.Scan(&document.Id,
			&document.DocumentId,
			&document.UserId,
			&document.RiskName,
			&document.FraudSchema,
			&document.FraudMotive,
			&document.FraudTechnique,
			&document.RiskSource,
			&document.RootCause,
			&document.BisproControlProcedure,
			&document.QualitativeImpact,
			&document.LikehoodJustification,
			&document.ImpactJustification,
			&document.StartegyAgreement,
			&document.StrategyRecomendation,
			&document.AssessmentLikehood,
			&document.AssessmentImpact,
			&document.AssessmentRiskLevel,
			&document.Action,
			&document.ActionBy,
			&document.CreatedAt,
			&document.UpdatedAt,
			&totalDocument,
		)
		if err != nil {
			return nil, 0, err
		}
		documents = append(documents, document)
	}

	return documents, totalDocument, nil
}
