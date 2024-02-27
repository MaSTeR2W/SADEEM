CREATE TABLE user_classifications(
    user_id INTEGER,
    class_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    FOREIGN KEY (class_id) REFERENCES classifications (class_id) ON DELETE CASCADE
);