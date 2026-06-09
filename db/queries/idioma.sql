-- name: GetAllIdioma :many
SELECT * FROM Idioma;

-- name: GetIdiomaById :one
SELECT * FROM Idioma WHERE idIdioma = ?;

-- name: CreateIdioma :execresult
INSERT INTO Idioma (nombre, fechaCreacion, fechaActualizacion)
VALUES (?, now(), now());

-- name: UpdateIdioma :execresult
UPDATE Idioma 
SET nombre = ?, 
    fechaActualizacion = now()
WHERE idIdioma = ?;

-- name: DeleteIdioma :execresult
DELETE FROM Idioma WHERE idIdioma = ?;