-- name: GetIdiomasByGuia :many
SELECT ig.ididiomaguia, ig.idguia, ig.ididioma, i.nombre AS nombreIdioma
FROM IdiomaGuia ig
JOIN Idioma i ON ig.idIdioma = i.idIdioma
WHERE ig.idGuia = ?;

-- name: CreateIdiomaGuia :execresult
INSERT INTO IdiomaGuia (idGuia, idIdioma, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());

-- name: DeleteIdiomaGuia :execresult
DELETE FROM IdiomaGuia WHERE idIdiomaGuia = ?;