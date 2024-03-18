CREATE TABLE users_history_login (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  INDEX users_history_login_idx_created_at (created_at),
  CONSTRAINT users_history_login_fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)