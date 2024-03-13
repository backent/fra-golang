CREATE TABLE documents_tracker (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  document_uuid VARCHAR(40) NOT NULL,
  viewed_count INT NOT NULL DEFAULT 0,
  searched_count INT NOT NULL DEFAULT 0,
  document_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,


  INDEX documents_tracker_document_uuid (document_uuid),
  INDEX documents_tracker_document_created_at (document_created_at),
  UNIQUE documents_tracker_u_docment_uuid (document_uuid)
)