TRUNCATE TABLE risks;

ALTER TABLE risks
DROP FOREIGN KEY risks_fk_document_id;

ALTER TABLE risks
MODIFY COLUMN document_id VARCHAR(40) NOT NULL;

ALTER TABLE risks
ADD UNIQUE(document_id);