CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(320),
    image VARCHAR(60),
    password BYTEA,
    salt BYTEA,
    user_type VARCHAR(10)
);