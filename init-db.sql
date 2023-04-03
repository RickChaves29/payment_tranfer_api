CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY,
        user_name VARCHAR(200) NOT NULL,
        email VARCHAR(150) UNIQUE NOT NULL,
        user_password VARCHAR(150) NOT NULL,
		balance BIGINT DEFAULT 0
	);
CREATE TABLE IF NOT EXISTS transfers (
    id SERIAL PRIMARY KEY,
    payer VARCHAR(200) NOT NULL,
    receive VARCHAR(200) NOT NULL,
    amount BIGINT NOT NULL,
    transfer_date TIMESTAMP DEFAULT NOW(),
    user_id BIGINT,
    CONSTRAINT fk_transfer_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

INSERT INTO users (user_name, email, user_password, balance) VALUES ('João sem maria', 'joao1234@gmail.com', '12345', 3000) ON CONFLICT (id) DO NOTHING;
INSERT INTO users (user_name, email, user_password, balance) VALUES ('João e maria', 'joao4321@gmail.com', '54321', 6000) ON CONFLICT (id) DO NOTHING;
