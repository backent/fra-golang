ALTER TABLE documents
DROP INDEX documents_uuid,
ADD CONSTRAINT documents_un_uuid UNIQUE (uuid)