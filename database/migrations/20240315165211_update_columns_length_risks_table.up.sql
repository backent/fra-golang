ALTER TABLE risks

MODIFY COLUMN fraud_schema TEXT NOT NULL,
MODIFY COLUMN fraud_motive TEXT NOT NULL,
MODIFY COLUMN fraud_technique TEXT NOT NULL,
MODIFY COLUMN risk_source TEXT NOT NULL,
MODIFY COLUMN root_cause TEXT NOT NULL,
MODIFY COLUMN bispro_control_procedure TEXT NOT NULL,
MODIFY COLUMN qualitative_impact TEXT NOT NULL,
MODIFY COLUMN likehood_justification TEXT NOT NULL,
MODIFY COLUMN impact_justification TEXT NOT NULL,
MODIFY COLUMN strategy_agreement TEXT NOT NULL,
MODIFY COLUMN strategy_recomendation TEXT NOT NULL