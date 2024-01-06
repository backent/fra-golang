ALTER TABLE documents
ADD COLUMN assessment_likehood VARCHAR(20) NOT NULL AFTER strategy_recomendation,
ADD COLUMN assessment_impact VARCHAR(20) NOT NULL AFTER assessment_likehood,
ADD COLUMN assessment_risk_level VARCHAR(20) NOT NULL AFTER assessment_impact