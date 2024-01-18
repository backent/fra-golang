ALTER TABLE `risks` DROP INDEX `document_id`;

TRUNCATE TABLE risks;

ALTER TABLE risks
MODIFY COLUMN document_id INT NOT NULL;

ALTER TABLE risks
ADD CONSTRAINT risks_fk_document_id
FOREIGN KEY (document_id)
REFERENCES documents(id) ON DELETE CASCADE;
