CREATE TABLE IF NOT EXISTS apikey (
    uid SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    timeCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    apiKey VARCHAR NOT NULL
);
