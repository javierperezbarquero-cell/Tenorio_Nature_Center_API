-- name: GetAllChofer :many
SELECT * FROM chofer;

-- name: GetChoferById :one
SELECT * FROM chofer WHERE idChofer = ?;

-- name: CreateChofer :execresult
INSERT INTO chofer (nombre, identificador, fechaNac, telefono, email, tipoLicencia, nacionalidad, fechaCreacion, fechaActualizacion)
VALUES (?, ?, ?, ?, ?, ?, ?, now(), now());

-- name: UpdateChofer :execresult
UPDATE chofer 
SET nombre = ?, 
    identificador = ?,
    fechaNac = ?, 
    telefono = ?, 
    email = ?, 
    tipoLicencia = ?, 
    nacionalidad = ?,
    fechaActualizacion = now()
WHERE idChofer = ?;

-- name: DeleteChofer :execresult
DELETE FROM chofer WHERE idChofer = ?;