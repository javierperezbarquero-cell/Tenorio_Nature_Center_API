-- name: GetAllIdiomaGuia :many
SELECT * FROM idiomaguia;

-- name: GetIdiomaGuiaById :one
SELECT * FROM idiomaguia WHERE idIdiomaGuia = ?;

-- name: CreateIdiomaGuia :execresult
INSERT INTO idiomaguia (idGuia, idIdioma, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());

-- name: UpdateIdiomaGuia :execresult
UPDATE idiomaguia 
SET idGuia = ?, 
    idIdioma = ?,
    fechaActualizacion = now()
WHERE idIdiomaGuia = ?;

-- name: DeleteIdiomaGuia :execresult
DELETE FROM idiomaguia WHERE idIdiomaGuia = ?;