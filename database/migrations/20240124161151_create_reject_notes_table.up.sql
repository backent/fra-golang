CREATE TABLE reject_notes (
  id INT NOT NULL AUTO_INCREMENT,
  document_id INT NOT NULL,
  risk_id INT NOT NULL,
  fraud VARCHAR(1000),
  risk_source VARCHAR(1000),
  root_cause VARCHAR(1000),
  bispro_control_procedure VARCHAR(1000),
  qualitative_impact VARCHAR(1000),
  assessment VARCHAR(1000),
  justification VARCHAR(1000),
  strategy VARCHAR(1000),

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY(id),
  CONSTRAINT reject_notes_fk_document_id FOREIGN KEY (document_id) REFERENCES documents(id) ON DELETE CASCADE,
  CONSTRAINT reject_notes_fk_risk_id FOREIGN KEY (risk_id) REFERENCES risks(id) ON DELETE CASCADE

)