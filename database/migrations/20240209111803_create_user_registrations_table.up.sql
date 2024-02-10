CREATE TABLE user_registrations (
  id INT NOT NULL AUTO_INCREMENT,
  nik VARCHAR(20) NOT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'pending',
  reject_by INT,
  approve_by INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX user_registrations_status (status)
)
