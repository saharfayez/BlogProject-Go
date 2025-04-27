CREATE TABLE app.role_permissions
(
    "role_id" INT,
    "permission_id" INT,
    PRIMARY KEY (role_id,permission_id),
    FOREIGN KEY (role_id) REFERENCES app.roles(id),
    FOREIGN KEY (permission_id) REFERENCES app.permissions(id)
);
