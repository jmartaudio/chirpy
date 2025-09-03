-- name: GetAllChirps :many
SELECT
    id,
    created_at,
    updated_at,
    body,
    user_id
FROM chirps
ORDER BY created_at ASC;

-- name: GetAllChipsByAuthor :many
SELECT
    id,
    created_at,
    updated_at,
    body,
    user_id
FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC;
