CREATE TABLE notifications (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  document_id INT NOT NULL,
  title VARCHAR(50),
  subtitle VARCHAR(100),
  `read` TINYINT NOT NULL DEFAULT 0,
  `action` VARCHAR(20) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  CONSTRAINT notifications_fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT notifications_fk_document_id FOREIGN KEY (document_id) REFERENCES documents(id) ON DELETE CASCADE

)