CREATE TABLE posts (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       title TEXT NOT NULL,
                       content TEXT NOT NULL,
                       author_id INTEGER NOT NULL,
                       FOREIGN KEY (author_id) REFERENCES users (id)
);
