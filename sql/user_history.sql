create table user_histories (
    id SERIAL PRIMARY KEY,
    userId INT, 
    text Text,
    rate FLOAT,
    pitch FLOAT,
    volume FLOAT,
    voiceName VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
)

