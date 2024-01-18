TRUNCATE TABLE risks;

ALTER TABLE risks
ADD COLUMN status VARCHAR(20) NULL DEFAULT NULL AFTER risk_name,
ADD user_id INT NOT NULL AFTER document_id,
ADD COLUMN action_by INT NOT NULL,
ADD COLUMN action VARCHAR(20) NULL DEFAULT NULL,

ADD INDEX documents_action_by (action_by),
ADD INDEX documents_action (action),

ADD CONSTRAINT documents_fk_action_by FOREIGN KEY (action_by) REFERENCES users(id),
ADD FOREIGN KEY (user_id) REFERENCES users(id)