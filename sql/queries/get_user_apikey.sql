-- name: GetUserByAPIKey :one
SELECT * FROM USERS
WHERE apikey=$1;