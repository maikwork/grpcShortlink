CREATE DATABASE docker;
USE docker;
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    long VARCHAR(10),
    short VARCHAR(10)
)
