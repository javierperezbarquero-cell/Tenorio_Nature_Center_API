-- name: GetAllGuia :many
SELECT * FROM Guia;

-- name: GetGuiaById :one
SELECT * FROM Guia WHERE idGuia = ?;

-- name: CreateGuia :execresult
INSERT INTO Guia (nombre, fechaNac, telefono, nacionalidad, email, fechaCreacion, fechaActualizacion)
VALUES (?, ?, ?, ?, ?, now(), now());

-- name: UpdateGuia :execresult
UPDATE Guia 
SET nombre = ?,  
    fechaNac = ?,  
    telefono = ?,
    nacionalidad = ?,
    email = ?,
    fechaCreacion = ?,
    fechaActualizacion = now()
WHERE idGuia = ?;

-- name: DeleteGuia :execresult
DELETE FROM Guia WHERE idGuia = ?;