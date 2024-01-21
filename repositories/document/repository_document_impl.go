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
		uuid,
		created_by,
		action_by,
		action,
		product_name
		) VALUES (?, ?, ?, ?, ?) `, models.DocumentTable)
	result, err := tx.ExecContext(ctx, query,
		document.Uuid,
		document.CreatedBy,
		document.ActionBy,
		document.Action,
		document.ProductName,
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
		uuid = ?,
		created_by = ?,
		action_by = ?,
		action = ?,
		product_name = ?
	WHERE id = ?`, models.DocumentTable)
	_, err := tx.ExecContext(ctx, query,

		document.Uuid,
		document.CreatedBy,
		document.ActionBy,
		document.Action,
		document.ProductName,

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

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s WHERE id = ?
		)
		SELECT 
		a.id,
		a.uuid,
		a.created_by,
		a.action_by,
		a.action,
		a.product_name,
		a.created_at,
		a.updated_at,
		b.id,
		b.name,
		b.nik,
		c.id,
		c.document_id,
		c.risk_name,
		c.fraud_schema,
		c.fraud_motive,
		c.fraud_technique,
		c.risk_source,
		c.root_cause,
		c.bispro_control_procedure,
		c.qualitative_impact,
		c.likehood_justification,
		c.impact_justification,
		c.strategy_agreement,
		c.strategy_recomendation,
		c.assessment_likehood,
		c.assessment_impact,
		c.assessment_risk_level,
		c.created_at,
		c.updated_at
	FROM (SELECT * FROM main_table) a
	LEFT JOIN %s b ON a.created_by = b.id
	LEFT JOIN %s c ON a.id = c.document_id `, models.DocumentTable, models.UserTable, models.RiskTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return document, err
	}
	defer rows.Close()

	var documents []*models.Document
	documentsMap := make(map[int]*models.Document)

	for rows.Next() {
		var document models.Document
		var user models.User
		var nullAbleRisk models.NullAbleRisk

		err = rows.Scan(
			&document.Id,
			&document.Uuid,
			&document.CreatedBy,
			&document.ActionBy,
			&document.Action,
			&document.ProductName,
			&document.CreatedAt,
			&document.UpdatedAt,
			&user.Id,
			&user.Name,
			&user.Nik,
			&nullAbleRisk.Id,
			&nullAbleRisk.DocumentId,
			&nullAbleRisk.RiskName,
			&nullAbleRisk.FraudSchema,
			&nullAbleRisk.FraudMotive,
			&nullAbleRisk.FraudTechnique,
			&nullAbleRisk.RiskSource,
			&nullAbleRisk.RootCause,
			&nullAbleRisk.BisproControlProcedure,
			&nullAbleRisk.QualitativeImpact,
			&nullAbleRisk.LikehoodJustification,
			&nullAbleRisk.ImpactJustification,
			&nullAbleRisk.StartegyAgreement,
			&nullAbleRisk.StrategyRecomendation,
			&nullAbleRisk.AssessmentLikehood,
			&nullAbleRisk.AssessmentImpact,
			&nullAbleRisk.AssessmentRiskLevel,
			&nullAbleRisk.CreatedAt,
			&nullAbleRisk.UpdatedAt,
		)
		if err != nil {
			return document, err
		}

		item, found := documentsMap[document.Id]
		if !found {
			item = &document
			documentsMap[document.Id] = item
			documents = append(documents, item)
		}
		item.UserDetail = user
		if nullAbleRisk.Id.Valid {
			item.RiskDetail = append(item.RiskDetail, models.NullAbleRiskToRisk(nullAbleRisk))
		}
	}

	if len(documents) < 1 {
		return document, errors.New("not found document")
	}

	var documentsReturn []models.Document
	for _, document := range documents {
		documentsReturn = append(documentsReturn, *document)
	}

	return documentsReturn[0], nil
}

func (implementation *RepositoryDocumentImpl) FindByUUID(ctx context.Context, tx *sql.Tx, documentUuid string) (models.Document, error) {
	var document models.Document
	query := fmt.Sprintf(`SELECT 
	id,
	uuid,
	created_by,
	action_by,
	action,
	product_name,
	created_at,
	updated_at
	FROM %s WHERE uuid = ?`, models.DocumentTable)
	rows, err := tx.QueryContext(ctx, query, documentUuid)
	if err != nil {
		return document, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&document.Id,
			&document.Id,
			&document.Uuid,
			&document.CreatedAt,
			&document.ActionBy,
			&document.Action,
			&document.ProductName,
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

func (implementation *RepositoryDocumentImpl) FindAllWithDetail(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Document, int, error) {
	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s
		)
		SELECT 
		a.id,
		a.uuid,
		a.created_by,
		a.action_by,
		a.action,
		a.product_name,
		a.created_at,
		a.updated_at,
		b.id,
		b.name,
		b.nik,
		c.id,
		c.document_id,
		c.risk_name,
		c.fraud_schema,
		c.fraud_motive,
		c.fraud_technique,
		c.risk_source,
		c.root_cause,
		c.bispro_control_procedure,
		c.qualitative_impact,
		c.likehood_justification,
		c.impact_justification,
		c.strategy_agreement,
		c.strategy_recomendation,
		c.assessment_likehood,
		c.assessment_impact,
		c.assessment_risk_level,
		c.created_at,
		c.updated_at,
		z.count
	FROM (SELECT * FROM main_table ORDER BY %s %s  LIMIT ?, ?) a
	LEFT JOIN %s b ON a.created_by = b.id
	LEFT JOIN %s c ON a.id = c.document_id
	LEFT JOIN (SELECT COUNT(*) as count FROM main_table) z ON true `, models.DocumentTable, orderBy, orderDirection, models.UserTable, models.RiskTable)
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
		var nullAbleRisk models.NullAbleRisk

		err = rows.Scan(
			&document.Id,
			&document.Uuid,
			&document.CreatedBy,
			&document.ActionBy,
			&document.Action,
			&document.ProductName,
			&document.CreatedAt,
			&document.UpdatedAt,
			&user.Id,
			&user.Name,
			&user.Nik,
			&nullAbleRisk.Id,
			&nullAbleRisk.DocumentId,
			&nullAbleRisk.RiskName,
			&nullAbleRisk.FraudSchema,
			&nullAbleRisk.FraudMotive,
			&nullAbleRisk.FraudTechnique,
			&nullAbleRisk.RiskSource,
			&nullAbleRisk.RootCause,
			&nullAbleRisk.BisproControlProcedure,
			&nullAbleRisk.QualitativeImpact,
			&nullAbleRisk.LikehoodJustification,
			&nullAbleRisk.ImpactJustification,
			&nullAbleRisk.StartegyAgreement,
			&nullAbleRisk.StrategyRecomendation,
			&nullAbleRisk.AssessmentLikehood,
			&nullAbleRisk.AssessmentImpact,
			&nullAbleRisk.AssessmentRiskLevel,
			&nullAbleRisk.CreatedAt,
			&nullAbleRisk.UpdatedAt,
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
		if nullAbleRisk.Id.Valid {
			item.RiskDetail = append(item.RiskDetail, models.NullAbleRiskToRisk(nullAbleRisk))
		}
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
		a.uuid,
		a.created_by,
		a.action_by,
		a.action,
		a.product_name,
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
		err = rows.Scan(
			&document.Id,
			&document.Uuid,
			&document.CreatedBy,
			&document.ActionBy,
			&document.Action,
			&document.ProductName,
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

func (implementation *RepositoryDocumentImpl) GetProductDistinct(ctx context.Context, tx *sql.Tx) ([]models.Document, error) {
	query := fmt.Sprintf(`
	WITH main_table AS (
		SELECT * FROM %s
	)
	SELECT d1.id, d1.product_name
	FROM main_table d1
	JOIN (
			SELECT product_name, MAX(id) AS max_id
			FROM main_table
			GROUP BY product_name
	) d2 ON d1.product_name = d2.product_name AND d1.id = d2.max_id`, models.DocumentTable)

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []models.Document
	for rows.Next() {
		var document models.Document
		err := rows.Scan(
			&document.Id,
			&document.ProductName,
		)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	return documents, nil

}
