-- name: GetAd :one
SELECT * FROM ads
WHERE id = $1 LIMIT 1;

-- name: ListAds :many
SELECT * FROM ads
ORDER BY created_at;

-- name: CreateAd :one
INSERT INTO ads (
  title, content,user_id,created_at,updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteAd :exec
DELETE FROM ads
WHERE id = $1;


-- name: UpdateAd :exec
UPDATE ads SET title = $2
WHERE id = $1;

-- name: GetAdsByUserID :many
SELECT * FROM ads
WHERE user_id = $1
ORDER BY created_at;
