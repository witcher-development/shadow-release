-- name: GetApp :one
select * from app
where id = ?;


-- name: CreateApp :one
INSERT INTO app (id) VALUES (?) RETURNING *;

-- name: GetVersion :one
select * from version
where name = ?;

-- name: CreateVersion :one
INSERT INTO version (name, app) VALUES (?, ?) RETURNING *;


-- name: CreateRecord :one
INSERT INTO record (version, path, reqbody, resbody) VALUES (?, ?, ?, ?) RETURNING *;