ALTER TABLE users
ADD COLUMN email VARCHAR(100) AFTER `name`,
ADD COLUMN apply_reject_by INT AFTER `password`,
ADD COLUMN apply_approved_by INT AFTER apply_reject_by,
ADD COLUMN apply_status VARCHAR(20) AFTER apply_approved_by,
MODIFY COLUMN password VARCHAR(200)