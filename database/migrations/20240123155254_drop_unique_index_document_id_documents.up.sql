ALTER TABLE documents
DROP INDEX documents_un_uuid,
ADD INDEX documents_uuid (uuid)