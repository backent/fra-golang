ALTER TABLE users
ADD COLUMN role VARCHAR(10) NOT NULL DEFAULT 'author' AFTER password