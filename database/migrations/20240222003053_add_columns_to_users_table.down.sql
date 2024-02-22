ALTER TABLE users
DROP COLUMN email,
DROP COLUMN apply_reject_by,
DROP COLUMN apply_approved_by,
DROP COLUMN apply_status,
MODIFY COLUMN password VARCHAR(200) NOT NULL