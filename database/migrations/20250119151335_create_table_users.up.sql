CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       user_name TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL
);
