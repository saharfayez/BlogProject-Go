INSERT INTO app.users
(id, username, password)
VALUES (1, 'admin', '$2a$10$cHYD6FwGuAnWWi8vlZUI/.Rfo2X8XIv5gmec.MYXz.F38vaWQdlme'),
(2, 'user', '$2a$10$cHYD6FwGuAnWWi8vlZUI/.Rfo2X8XIv5gmec.MYXz.F38vaWQdlme');

INSERT INTO app.roles
(id, name)
VALUES (1, 'admin'),
(2, 'user');

INSERT INTO app.permissions
(id, name)
VALUES (1, 'POST'),
(2, 'GET'),
(3, 'PUT'),
(4, 'DELETE');

INSERT INTO app.resources
(id, name)
VALUES (1, '/api/posts');

INSERT INTO app.role_permissions
(role_id, permission_id)
VALUES (1, 1),
(1, 2),
(1, 3),
(1, 4),
(2, 2);

INSERT INTO app.user_roles
(user_id, role_id)
VALUES (1, 1),
(2, 2);

INSERT INTO app.resource_permissions
(permission_id, resource_id)
VALUES (1, 1);
