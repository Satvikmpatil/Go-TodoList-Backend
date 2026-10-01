-- name: CreateTask :one
INSERT INTO tasks (user_id, task)
VALUES ($1, $2)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE task_id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY created_at DESC;

-- name: ListTasksByUser :many
SELECT * FROM tasks
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateTaskStatus :one
UPDATE tasks
SET status = $2
WHERE task_id = $1
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET task = $2, status = $3
WHERE task_id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE task_id = $1;
