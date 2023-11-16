-- name: CreateAyah :one

INSERT INTO
    ayah (
        page,
        surah_number,
        number,
        number_in_surah,
        ar,
        tr,
        idn,
        audio_url
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetAyah :one

SELECT * FROM ayah WHERE id = $1 LIMIT 1;

-- name: GetAyahForUpdate :one

SELECT * FROM ayah WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: ListAllAyah :many

SELECT * FROM ayah ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListAyahByPage :many

SELECT * FROM ayah WHERE page = $1;

-- name: UpdateAyah :one

UPDATE ayah
SET
    page = $2,
    surah_number = $3,
    number = $4,
    ar = $5,
    tr = $6,
    idn = $7,
    audio_url = $8,
    number_in_surah = $9
WHERE id = $1 RETURNING *;

-- name: DeleteAyah :exec

DELETE FROM ayah WHERE id = $1;