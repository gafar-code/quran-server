-- name: CreateSurah :one

INSERT INTO
    surah (
        page,
        ar,
        tr,
        idn,
        number,
        type,
        total_ayah
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetSurah :one

SELECT * FROM surah WHERE id = $1 LIMIT 1;
-- name: GetSurahByNumber :one

SELECT * FROM surah WHERE number = $1 LIMIT 1;

-- name: ListSurah :many

SELECT * FROM surah ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListSurahByPage :many

SELECT * FROM surah WHERE page = $1;

-- name: UpdateSurah :one

UPDATE surah
SET
    page = $2,
    ar = $3,
    tr = $4,
    idn = $5,
    number = $6,
    type = $7,
    total_ayah = $8
WHERE id = $1 RETURNING *;

-- name: DeleteSurah :exec

DELETE FROM surah WHERE id = $1;