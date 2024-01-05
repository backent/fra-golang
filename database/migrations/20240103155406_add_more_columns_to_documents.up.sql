ALTER TABLE documents
ADD COLUMN status VARCHAR(20) NULL DEFAULT NULL AFTER risk_name,
ADD COLUMN action_by INT NOT NULL,
ADD COLUMN action VARCHAR(20) NULL DEFAULT NULL,

ADD INDEX documents_risk_name (risk_name),
ADD INDEX documents_action_by (action_by),
ADD INDEX documents_action (action),

ADD CONSTRAINT documents_fk_action_by FOREIGN KEY (action_by) REFERENCES users(id)