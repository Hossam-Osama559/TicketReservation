ALTER TABLE sessions
ADD COLUMN userid INT,
ADD CONSTRAINT fk_sessions_user
FOREIGN KEY (userid) REFERENCES users(id);