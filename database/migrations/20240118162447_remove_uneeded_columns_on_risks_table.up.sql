ALTER TABLE risks
DROP FOREIGN KEY documents_fk_action_by,
DROP FOREIGN KEY risks_ibfk_1;

ALTER TABLE risks
DROP COLUMN user_id,
DROP COLUMN status,
DROP COLUMN action_by,
DROP COLUMN action;