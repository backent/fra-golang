ALTER TABLE documents
ADD COLUMN file_name VARCHAR(50) AFTER `category`,
ADD COLUMN file_original_name VARCHAR(200) AFTER `file_name`