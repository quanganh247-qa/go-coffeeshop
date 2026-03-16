BEGIN;

DROP TABLE IF EXISTS auth.refresh_tokens;
DROP TABLE IF EXISTS auth.role_permissions;
DROP TABLE IF EXISTS auth.user_roles;
DROP TABLE IF EXISTS auth.permissions;
DROP TABLE IF EXISTS auth.roles;
DROP TABLE IF EXISTS auth.users;

DROP SCHEMA IF EXISTS auth CASCADE;

COMMIT;