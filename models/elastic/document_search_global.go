package elastic

import (
	"strings"

	"github.com/backent/fra-golang/models"
)

const IndexNameDocumentSearchGlobal = "fra_document_search_global"
const IndexNameDocumentSearchGlobalSettings = `
{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "long"
			},
			"uuid": {
				"type": "keyword"
			},
			"action": {
				"type": "text"
			},
			"product_name": {
				"type": "text"
			},
			"category": {
				"type": "text"
			},
			"risk": {
				"type": "nested",
				"properties": {
					"risk_name": {
						"type": "text"
					},
					"fraud_motive": {
						"type": "text"
					},
					"fraud_technique": {
						"type": "text"
					},
					"risk_source": {
						"type": "text"
					},
					"root_cause": {
						"type": "text"
					},
					"bispro_control_procedure": {
						"type": "text"
					},
					"qualitative_impact": {
						"type": "text"
					},
					"likehood_justification": {
						"type": "text"
					},
					"impact_justification": {
						"type": "text"
					},
					"strategy_agreement": {
						"type": "keyword"
					},
					"strategy_recomendation": {
						"type": "text"
					},
					"assessment_likehood": {
						"type": "keyword"
					},
					"assessment_impact": {
						"type": "keyword"
					},
					"assessment_risk_level": {
						"type": "keyword"
					}
				}
			}
		}
	}
}`

type DocumentSearchGlobal struct {
	Id               int                `json:"id"`           // id
	Uuid             string             `json:"uuid"`         // uuid
	Action           string             `json:"action"`       // action
	ProductName      string             `json:"product_name"` // product_name
	Category         string             `json:"category"`     // category
	RiskSearchGlobal []RiskSearchGlobal `json:"risk"`         // risk
}

type RiskSearchGlobal struct {
	RiskName               string `json:"risk_name"`                // risk_name
	FraudMotive            string `json:"fraud_motive"`             // fraud_motive
	FraudTechnique         string `json:"fraud_technique"`          // fraud_technique
	FraudSchema            string `json:"fraud_schema"`             // fraud_technique
	RiskSource             string `json:"risk_source"`              // risk_source
	RootCause              string `json:"root_cause"`               // root_cause
	BisproControlProcedure string `json:"bispro_control_procedure"` // bispro_control_procedure
	QualitativeImpact      string `json:"qualitative_impact"`       // qualitative_impact
	LikehoodJustification  string `json:"likehood_justification"`   // likehood_justification
	ImpactJustification    string `json:"impact_justification"`     // impact_justification
	StartegyAgreement      string `json:"strategy_agreement"`       // strategy_agreement
	StrategyRecomendation  string `json:"strategy_recomendation"`   // strategy_recomendation
	AssessmentLikehood     string `json:"assessment_likehood"`      // assessment_likehood
	AssessmentImpact       string `json:"assessment_impact"`        // assessment_impact
	AssessmentRiskLevel    string `json:"assessment_risk_level"`    // assessment_risk_level
}

func ModelDocumentToIndexDocumentSearchGlobal(document models.Document) DocumentSearchGlobal {
	return DocumentSearchGlobal{
		Id:               document.Id,
		Uuid:             document.Uuid,
		ProductName:      document.ProductName,
		Action:           document.Action,
		Category:         document.Category,
		RiskSearchGlobal: BulkModelRiskToIndexRiskSearchGlobal(document.RiskDetail),
	}
}

func BulkModelDocumentToIndexDocumentSearchGlobal(documents []models.Document) []DocumentSearchGlobal {
	var bulk []DocumentSearchGlobal
	for _, val := range documents {
		bulk = append(bulk, ModelDocumentToIndexDocumentSearchGlobal(val))
	}

	return bulk
}

func ModelRiskToIndexRiskSearchGlobal(risk models.Risk) RiskSearchGlobal {
	return RiskSearchGlobal{
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
	}
}

func BulkModelRiskToIndexRiskSearchGlobal(risks []models.Risk) []RiskSearchGlobal {
	var bulk []RiskSearchGlobal
	for _, val := range risks {
		bulk = append(bulk, ModelRiskToIndexRiskSearchGlobal(val))
	}
	return bulk
}

func GenerateQuery(query string) string {
	return strings.ReplaceAll(`{
		"query": {
			"bool": {
				"should": [
					{
						"match": {
							"product_name": {
								"query": "{{value}}", 
								"fuzziness": "auto"
							}
						}
					},
					{
						"nested": {
							"path": "risk",
							"query": {
								"bool": {
									"should": [
										{
											"match": {
												"risk.risk_name": {
													"query": "{{value}}",
													"fuzziness": "auto"
												}
											}
										}
									]
								}
							}
						}
					}
				]
			}
		}
	}`, "{{value}}", query)
}
