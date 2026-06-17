-- name: GetAllChofer :many
SELECT * FROM Chofer;

-- name: GetChoferById :one
SELECT * FROM Chofer WHERE idChofer = ?;

-- name: CreateChofer :execresult
INSERT INTO Chofer (nombre, identificador, fechaNac, telefono, email, tipoLicencia, nacionalidad, fechaCreacion, fechaActualizacion)
VALUES (?, ?, ?, ?, ?, ?, ?, now(), now());

-- name: UpdateChofer :execresult
UPDATE Chofer 
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
DELETE FROM Chofer WHERE idChofer = ?;