CREATE TABLE tickets (
    id INT AUTO_INCREMENT PRIMARY KEY,

    origin VARCHAR(255) NOT NULL,
    arrival VARCHAR(255) NOT NULL,

    price INT NOT NULL,

    user_id INT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE

);