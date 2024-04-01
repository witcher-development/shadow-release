-- name: GetApp :one
select * from app
where id = ?;


-- name: CreateApp :one
INSERT INTO app (id) VALUES ( ? ) RETURNING *;
