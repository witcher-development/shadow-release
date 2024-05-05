-- name: GetApp :one
select * from app
where id = ?;

-- name: CreateApp :one
INSERT INTO app (id) VALUES (?) RETURNING *;

-- name: CreateVersion :one
INSERT INTO version (name, app) VALUES (?, ?) RETURNING *;


-- name: CreateRecord :one
INSERT INTO record (version, path, method, reqbody, resbody, synckey) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetRecords :many
select * from record;

-- name: GetVersions :many
select * from version;

-- name: GetVersion :one
select * from version
where name = ?;
