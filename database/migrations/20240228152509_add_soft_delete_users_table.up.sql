ALTER TABLE users
ADD COLUMN delete_by INT AFTER `role`,
ADD COLUMN deleted_at TIMESTAMP