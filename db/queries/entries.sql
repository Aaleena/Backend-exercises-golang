-- name: AddEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :many
SELECT * FROM entries
WHERE account_id = $1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries 
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE from entries WHERE id = $1;