CREATE TABLE app.resource_permissions
(
    "permission_id" INT,
    "resource_id"   INT,
    PRIMARY KEY (permission_id, resource_id),
    FOREIGN KEY (permission_id) REFERENCES app.permissions (id),
    FOREIGN KEY (resource_id) REFERENCES app.resources (id)
);
