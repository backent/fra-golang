package risk

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/backent/fra-golang/models"
)

type RepositoryRiskImpl struct {
}

func NewRepositoryRiskImpl() RepositoryRiskInterface {
	return &RepositoryRiskImpl{}
}

func (implementation *RepositoryRiskImpl) Create(ctx context.Context, tx *sql.Tx, risk models.Risk) (models.Risk, error) {
	query := fmt.Sprintf(`INSERT INTO %s (
		document_id,
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
		assessment_risk_level
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `, models.RiskTable)
	result, err := tx.ExecContext(ctx, query,
		risk.DocumentId,
		risk.RiskName,
		risk.FraudSchema,
		risk.FraudMotive,
		risk.FraudTechnique,
		risk.RiskSource,
		risk.RootCause,
		risk.BisproControlProcedure,
		risk.QualitativeImpact,
		risk.LikehoodJustification,
		risk.ImpactJustification,
		risk.StartegyAgreement,
		risk.StrategyRecomendation,
		risk.AssessmentLikehood,
		risk.AssessmentImpact,
		risk.AssessmentRiskLevel,
	)
	if err != nil {
		return risk, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return risk, err
	}

	risk.Id = int(id)

	return risk, nil

}
func (implementation *RepositoryRiskImpl) Update(ctx context.Context, tx *sql.Tx, risk models.Risk) (models.Risk, error) {
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
		assessment_risk_level = ?
	WHERE id = ?`, models.RiskTable)
	_, err := tx.ExecContext(ctx, query,

		risk.RiskName,
		risk.FraudSchema,
		risk.FraudMotive,
		risk.FraudTechnique,
		risk.RiskSource,
		risk.RootCause,
		risk.BisproControlProcedure,
		risk.QualitativeImpact,
		risk.LikehoodJustification,
		risk.ImpactJustification,
		risk.StartegyAgreement,
		risk.StrategyRecomendation,
		risk.AssessmentLikehood,
		risk.AssessmentImpact,
		risk.AssessmentRiskLevel,

		risk.Id)
	if err != nil {
		return risk, err
	}

	return risk, nil
}
func (implementation *RepositoryRiskImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := fmt.Sprintf("DELETE FROM  %s  WHERE id = ?", models.RiskTable)
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (implementation *RepositoryRiskImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.Risk, error) {
	var risk models.Risk

	query := fmt.Sprintf(`SELECT 
		id,
		document_id,
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
		created_at,
		updated_at
	FROM %s WHERE id = ?`, models.RiskTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return risk, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&risk.Id,
			&risk.DocumentId,
			&risk.RiskName,
			&risk.FraudSchema,
			&risk.FraudMotive,
			&risk.FraudTechnique,
			&risk.RiskSource,
			&risk.RootCause,
			&risk.BisproControlProcedure,
			&risk.QualitativeImpact,
			&risk.LikehoodJustification,
			&risk.ImpactJustification,
			&risk.StartegyAgreement,
			&risk.StrategyRecomendation,
			&risk.AssessmentLikehood,
			&risk.AssessmentImpact,
			&risk.AssessmentRiskLevel,
			&risk.CreatedAt,
			&risk.UpdatedAt,
		)
		if err != nil {
			return risk, err
		}
	} else {
		return risk, errors.New("not found risk")
	}

	return risk, nil
}

func (implementation *RepositoryRiskImpl) FindByDocumentId(ctx context.Context, tx *sql.Tx, riskId string) (models.Risk, error) {
	var risk models.Risk
	query := fmt.Sprintf(`SELECT 
		id,
		document_id,
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
		created_at,
		updated_at
	FROM %s WHERE document_id = ?`, models.RiskTable)
	rows, err := tx.QueryContext(ctx, query, riskId)
	if err != nil {
		return risk, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&risk.Id,
			&risk.DocumentId,
			&risk.RiskName,
			&risk.FraudSchema,
			&risk.FraudMotive,
			&risk.FraudTechnique,
			&risk.RiskSource,
			&risk.RootCause,
			&risk.BisproControlProcedure,
			&risk.QualitativeImpact,
			&risk.LikehoodJustification,
			&risk.ImpactJustification,
			&risk.StartegyAgreement,
			&risk.StrategyRecomendation,
			&risk.AssessmentLikehood,
			&risk.AssessmentImpact,
			&risk.AssessmentRiskLevel,
			&risk.CreatedAt,
			&risk.UpdatedAt,
		)
		if err != nil {
			return risk, err
		}
	} else {
		return risk, errors.New("not found risk")
	}

	return risk, nil
}

func (implementation *RepositoryRiskImpl) FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Risk, int, error) {
	var risks []models.Risk

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s
		)
		SELECT 
		a.id,
		a.document_id,
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
		a.created_at,
		a.updated_at,
		b.count
	FROM (SELECT * FROM main_table ORDER BY %s %s LIMIT ?, ?) a LEFT JOIN (SELECT COUNT(*) as count FROM main_table) b ON true`, models.RiskTable, orderBy, orderDirection)
	rows, err := tx.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var totalRisk int
	for rows.Next() {
		var risk models.Risk
		err = rows.Scan(&risk.Id,
			&risk.DocumentId,
			&risk.RiskName,
			&risk.FraudSchema,
			&risk.FraudMotive,
			&risk.FraudTechnique,
			&risk.RiskSource,
			&risk.RootCause,
			&risk.BisproControlProcedure,
			&risk.QualitativeImpact,
			&risk.LikehoodJustification,
			&risk.ImpactJustification,
			&risk.StartegyAgreement,
			&risk.StrategyRecomendation,
			&risk.AssessmentLikehood,
			&risk.AssessmentImpact,
			&risk.AssessmentRiskLevel,
			&risk.CreatedAt,
			&risk.UpdatedAt,
			&totalRisk,
		)
		if err != nil {
			return nil, 0, err
		}
		risks = append(risks, risk)
	}

	return risks, totalRisk, nil
}
