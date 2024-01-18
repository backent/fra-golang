CREATE TABLE documents (
  id INT NOT NULL AUTO_INCREMENT,
  uuid VARCHAR(40) NOT NULL,
  created_by INT NOT NULL,
  action_by INT NOT NULL,
  action VARCHAR(20) NOT NULL,
  product_name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  CONSTRAINT documents_fk_created_by FOREIGN KEY (created_by) REFERENCES users(id),
  CONSTRAINT documents_un_uuid UNIQUE (uuid),
  INDEX documents_product_name (product_name),
  INDEX documents_created_by (created_by),
  INDEX documents_action_by (action_by),
  INDEX documents_action (action)
)