CREATE TABLE app.roles
(
    "id"       SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE
);
