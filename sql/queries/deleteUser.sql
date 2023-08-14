-- name: DeleteUser :exec
DELETE from users WHERE name = $1;