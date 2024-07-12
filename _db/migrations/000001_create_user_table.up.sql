CREATE TABLE IF NOT EXISTS users (
                                     user_id SERIAL PRIMARY KEY,
                                     login TEXT NOT NULL,
                                     first_name TEXT NOT NULL,
                                     second_name TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     role_id INT NOT NULL,
                                     FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

CREATE INDEX idx_users_login ON users(login);
