-- name: GetIdiomasByGuia :many
SELECT ig.ididiomaguia, ig.idguia, ig.ididioma, i.nombre AS nombreIdioma
-- name: GetAllIdiomaGuia :many
SELECT
    ig.idIdiomaGuia,
    ig.idGuia,
    ig.idIdioma,
    ig.fechaCreacion,
    ig.fechaActualizacion,

    -- Guia
    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,
    g.email     AS guiaEmail,

    -- Idioma
    i.nombre    AS idiomaNombre

FROM IdiomaGuia ig

JOIN Guia g     ON ig.idGuia   = g.idGuia
JOIN Idioma i   ON ig.idIdioma = i.idIdioma;


-- name: GetIdiomaGuiaById :one
SELECT
    ig.idIdiomaGuia,
    ig.idGuia,
    ig.idIdioma,
    ig.fechaCreacion,
    ig.fechaActualizacion,

    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,
    g.email     AS guiaEmail,

    i.nombre    AS idiomaNombre

FROM IdiomaGuia ig

JOIN Guia g     ON ig.idGuia   = g.idGuia
JOIN Idioma i   ON ig.idIdioma = i.idIdioma

WHERE ig.idIdiomaGuia = ?;


-- name: GetIdiomaGuiaByGuia :many
SELECT
    ig.idIdiomaGuia,
    ig.idGuia,
    ig.idIdioma,
    ig.fechaCreacion,
    ig.fechaActualizacion,

    -- Solo Idioma, el guía ya se conoce por el filtro
    i.nombre    AS idiomaNombre

FROM IdiomaGuia ig

JOIN Idioma i ON ig.idIdioma = i.idIdioma

WHERE ig.idGuia = ?;


-- name: GetIdiomaGuiaByIdioma :many
SELECT
    ig.idIdiomaGuia,
    ig.idGuia,
    ig.idIdioma,
    ig.fechaCreacion,
    ig.fechaActualizacion,

    -- Solo Guia, el idioma ya se conoce por el filtro
    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,
    g.email     AS guiaEmail

FROM IdiomaGuia ig

JOIN Guia g ON ig.idGuia = g.idGuia

WHERE ig.idIdioma = ?;


-- name: CreateIdiomaGuia :execresult
INSERT INTO IdiomaGuia (
    idGuia,
    idIdioma,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, now(), now());


-- name: DeleteIdiomaGuia :execresult
DELETE FROM IdiomaGuia WHERE idIdiomaGuia = ?;