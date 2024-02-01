package document

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/backent/fra-golang/helpers"
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
		),
		risk_with_reject_note AS (
			SELECT
			a.*,
			b.id as reject_note_id,
			b.fraud as reject_note_fraud,
			b.risk_source as reject_note_risk_source,
			b.root_cause as reject_note_root_cause,
			b.bispro_control_procedure as reject_note_bispro_control_procedure,
			b.qualitative_impact as reject_note_qualitative_impact,
			b.assessment as reject_note_assessment,
			b.justification as reject_note_justification,
			b.strategy as reject_note_strategy 
			FROM %s a LEFT JOIN %s b ON a.id = b.risk_id
		),
		related_document AS (
			SELECT * FROM %s WHERE action != 'draft' ORDER BY id DESC
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
		c.reject_note_id,
		c.reject_note_fraud,
		c.reject_note_risk_source,
		c.reject_note_root_cause,
		c.reject_note_bispro_control_procedure,
		c.reject_note_qualitative_impact,
		c.reject_note_assessment,
		c.reject_note_justification,
		c.reject_note_strategy,
		c.created_at,
		c.updated_at,
		d.id,
		d.created_at
	FROM (SELECT * FROM main_table) a
	LEFT JOIN %s b ON a.created_by = b.id
	LEFT JOIN risk_with_reject_note c ON a.id = c.document_id 
	LEFT JOIN related_document d ON a.uuid = d.uuid AND a.id != d.id
	ORDER BY d.id DESC, c.id ASC`, models.DocumentTable, models.RiskTable, models.RejectNoteTable, models.DocumentTable, models.UserTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return document, err
	}
	defer rows.Close()

	var documents []*models.Document
	documentsMap := make(map[int]*models.Document)
	documentWithRiskMap := make(map[string]bool)
	documentWithRelatedMap := make(map[string][]models.RelatedDocument)

	for rows.Next() {
		var document models.Document
		var user models.User
		var nullAbleRisk models.NullAbleRisk
		var nullAbleRejectNote models.NullAbleRejectNote
		var nullAbleRelatedDocument models.NullAbleRelatedDocument

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
			&nullAbleRejectNote.Id,
			&nullAbleRejectNote.Fraud,
			&nullAbleRejectNote.RiskSource,
			&nullAbleRejectNote.RootCause,
			&nullAbleRejectNote.BisproControlProcedure,
			&nullAbleRejectNote.QualitativeImpact,
			&nullAbleRejectNote.Assessment,
			&nullAbleRejectNote.Justification,
			&nullAbleRejectNote.Strategy,
			&nullAbleRisk.CreatedAt,
			&nullAbleRisk.UpdatedAt,
			&nullAbleRelatedDocument.Id,
			&nullAbleRelatedDocument.CreatedAt,
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
			riskDetail := models.NullAbleRiskToRisk(nullAbleRisk)
			if nullAbleRejectNote.Id.Valid {
				riskDetail.RejectNoteDetail.Id = int(nullAbleRejectNote.Id.Int32)
				riskDetail.RejectNoteDetail.DocumentId = item.Id
				riskDetail.RejectNoteDetail.RiskId = riskDetail.Id
				riskDetail.RejectNoteDetail.Fraud = nullAbleRejectNote.Fraud.String
				riskDetail.RejectNoteDetail.RiskSource = nullAbleRejectNote.RiskSource.String
				riskDetail.RejectNoteDetail.RootCause = nullAbleRejectNote.RootCause.String
				riskDetail.RejectNoteDetail.BisproControlProcedure = nullAbleRejectNote.BisproControlProcedure.String
				riskDetail.RejectNoteDetail.QualitativeImpact = nullAbleRejectNote.QualitativeImpact.String
				riskDetail.RejectNoteDetail.Assessment = nullAbleRejectNote.Assessment.String
				riskDetail.RejectNoteDetail.Justification = nullAbleRejectNote.Justification.String
				riskDetail.RejectNoteDetail.Strategy = nullAbleRejectNote.Strategy.String
			}
			if _, foundRisk := documentWithRiskMap[helpers.PrintStringIDRelation(document.Id, riskDetail.Id)]; !foundRisk {
				documentWithRiskMap[helpers.PrintStringIDRelation(document.Id, riskDetail.Id)] = true
				item.RiskDetail = append(item.RiskDetail, riskDetail)
			}
		}
		if nullAbleRelatedDocument.Id.Valid {
			if _, found := documentWithRelatedMap[helpers.PrintStringIDRelation(document.Id, int(nullAbleRelatedDocument.Id.Int32))]; !found {
				documentWithRelatedMap[helpers.PrintStringIDRelation(document.Id, int(nullAbleRelatedDocument.Id.Int32))] = append(documentWithRelatedMap[helpers.PrintStringIDRelation(document.Id, int(nullAbleRelatedDocument.Id.Int32))], models.NullAbleRelatedDocumentToRelatedDocument(nullAbleRelatedDocument))
				item.RelatedDocumentDetail = append(item.RelatedDocumentDetail, models.NullAbleRelatedDocumentToRelatedDocument(nullAbleRelatedDocument))
			}
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
		err = rows.Scan(
			&document.Id,
			&document.Uuid,
			&document.CreatedBy,
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

func (implementation *RepositoryDocumentImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
	take int,
	skip int,
	orderBy string,
	orderDirection string,
	userId int,
	documentAction string,
) ([]models.Document, int, error) {
	var documents []models.Document

	var conditionalQueryUser string
	if userId == 0 {
		userId = 1
		conditionalQueryUser = "AND 1 = ?"
	} else {
		conditionalQueryUser = "AND created_by = ?"
	}

	var conditionalQueryAction string
	var conditionalQueryValue []interface{}
	if documentAction == "" {
		documentAction = "1"
		conditionalQueryAction = "AND 1 = ?"
		conditionalQueryValue = append(conditionalQueryValue, "1")
	} else {
		for _, val := range strings.Split(documentAction, ",") {
			conditionalQueryValue = append(conditionalQueryValue, val)
		}
		helpers.Placeholders(len(conditionalQueryValue))
		conditionalQueryAction = fmt.Sprintf("AND action IN (%s)", helpers.Placeholders(len(conditionalQueryValue)))
	}

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s 
		), group_by_uuid AS (
			SELECT d1.*
			FROM main_table d1
			JOIN (
					SELECT uuid, MAX(id) AS max_id
					FROM main_table
					GROUP BY uuid
			) d2 ON d1.uuid = d2.uuid AND d1.id = d2.max_id
		), main_table_after_grouped AS (
			SELECT * FROM group_by_uuid WHERE 1 = 1 %s %s
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
	FROM (SELECT * FROM main_table_after_grouped ORDER BY %s %s LIMIT ?, ?) a LEFT JOIN (SELECT COUNT(*) as count FROM main_table_after_grouped) b ON true`, models.DocumentTable, conditionalQueryUser, conditionalQueryAction, orderBy, orderDirection)

	var args []interface{}
	args = append(args, userId)
	args = append(args, conditionalQueryValue...)
	args = append(args, skip, take)
	rows, err := tx.QueryContext(ctx, query, args...)
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

func (implementation *RepositoryDocumentImpl) FindAllNoGroup(
	ctx context.Context,
	tx *sql.Tx,
	take int,
	skip int,
	orderBy string,
	orderDirection string,
	documentAction string,
	month int,
	name string,
) ([]models.Document, int, error) {
	var documents []models.Document

	var conditionalQueryAction string
	var conditionalQueryValue []interface{}
	if documentAction == "" {
		documentAction = "1"
		conditionalQueryAction = "AND 1 = ?"
		conditionalQueryValue = append(conditionalQueryValue, "1")
	} else {
		for _, val := range strings.Split(documentAction, ",") {
			conditionalQueryValue = append(conditionalQueryValue, val)
		}
		helpers.Placeholders(len(conditionalQueryValue))
		conditionalQueryAction = fmt.Sprintf("AND action IN (%s)", helpers.Placeholders(len(conditionalQueryValue)))
	}

	var conditionalQueryPeriod string
	var conditionalQueryPeriodValue int
	if month == 0 {
		conditionalQueryPeriod = "AND 1 = ?"
		conditionalQueryPeriodValue = 1
	} else {
		conditionalQueryPeriod = "AND MONTH(created_at) = ?"
		conditionalQueryPeriodValue = month
	}

	var conditionalQueryName string
	var conditionalQueryNameValue string
	if name == "" {
		conditionalQueryName = "AND 1 = ?"
		conditionalQueryNameValue = "1"
	} else {
		conditionalQueryName = "AND product_name LIKE ?"
		conditionalQueryNameValue = "%" + name + "%"
	}

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s WHERE 1 = 1 AND YEAR(created_at) = YEAR(NOW()) %s %s %s
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
	FROM (SELECT * FROM main_table ORDER BY %s %s LIMIT ?, ?) a LEFT JOIN (SELECT COUNT(*) as count FROM main_table) b ON true`, models.DocumentTable, conditionalQueryAction, conditionalQueryPeriod, conditionalQueryName, orderBy, orderDirection)
	var args []interface{}
	args = append(args, conditionalQueryValue...)
	args = append(args, conditionalQueryPeriodValue)
	args = append(args, conditionalQueryNameValue)
	args = append(args, skip, take)
	rows, err := tx.QueryContext(ctx, query, args...)
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

func (implementation *RepositoryDocumentImpl) GetNonDraftProductByUUID(ctx context.Context, tx *sql.Tx, uuid string) ([]models.Document, error) {
	query := fmt.Sprintf(`
		SELECT
			id,
			product_name,
			uuid,
			action
		FROM %s WHERE action != 'draft' AND uuid = ?
	`, models.DocumentTable)

	rows, err := tx.QueryContext(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []models.Document
	for rows.Next() {
		var document models.Document
		err = rows.Scan(&document.Id, &document.ProductName, &document.Uuid, &document.Action)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	return documents, nil

}
