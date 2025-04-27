CREATE TABLE app.user_roles
(
    "user_id" INT,
    "role_id" INT,
    PRIMARY KEY (user_id,role_id),
    FOREIGN KEY (user_id) REFERENCES app.users(id),
    FOREIGN KEY (role_id) REFERENCES app.roles(id)
);
