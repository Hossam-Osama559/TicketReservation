ALTER TABLE sessions
DROP FOREIGN KEY fk_sessions_user,
DROP COLUMN userid;