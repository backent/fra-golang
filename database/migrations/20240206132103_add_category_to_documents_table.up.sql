ALTER TABLE documents
ADD COLUMN category VARCHAR(20) NOT NULL DEFAULT 'communication' AFTER product_name,

ADD INDEX documents_category (category)