-- name: GetAllUbicacion :many
SELECT * FROM Ubicacion;

-- name: GetUbicacionById :one
SELECT * FROM Ubicacion WHERE idUbicacion = ?;

-- name: CreateUbicacion :execresult
INSERT INTO Ubicacion (nombre, direccion, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());

-- name: UpdateUbicacion :execresult
UPDATE Ubicacion
SET nombre             = ?,
    direccion          = ?,
    fechaActualizacion = now()
WHERE idUbicacion = ?;

-- name: DeleteUbicacion :execresult
DELETE FROM Ubicacion WHERE idUbicacion = ?;