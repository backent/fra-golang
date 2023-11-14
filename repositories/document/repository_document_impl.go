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
		strategy_recomendation
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `, models.DocumentTable)
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
		strategy_recomendation = ?
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
		strategy_recomendation
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
		)
		if err != nil {
			return document, err
		}
	} else {
		return document, errors.New("not found document")
	}

	return document, nil
}
func (implementation *RepositoryDocumentImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]models.Document, error) {
	var documents []models.Document

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
		strategy_recomendation
	FROM %s ORDER BY id DESC`, models.DocumentTable)
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
		)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	return documents, nil
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
		strategy_recomendation
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
		)
		if err != nil {
			return document, err
		}
	} else {
		return document, errors.New("not found document")
	}

	return document, nil
}
