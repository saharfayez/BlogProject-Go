CREATE TABLE "users"
(
    "id"       SERIAL PRIMARY KEY,
    "username" VARCHAR(50),
    "password" VARCHAR(255),
    "role" VARCHAR(50)
);
