ALTER TABLE users
ADD COLUMN unit VARCHAR(20) AFTER `role`,

ADD INDEX users_idx_unit (unit)

