-- name: GetAllUbicacion :many
SELECT * FROM ubicacion;

-- name: GetUbicacionById :one
SELECT * FROM ubicacion WHERE idUbicacion = ?;

-- name: CreateUbicacion :execresult
INSERT INTO ubicacion (nombre, direccion, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());

-- name: UpdateUbicacion :execresult
UPDATE ubicacion
SET nombre                = ?,
    direccion              = ?,
    fechaActualizacion     = now()
WHERE idUbicacion = ?;

-- name: DeleteUbicacion :execresult
DELETE FROM ubicacion WHERE idUbicacion = ?;
