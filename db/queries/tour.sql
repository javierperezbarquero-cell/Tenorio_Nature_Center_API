-- name: CreateTour :execresult
INSERT INTO Tour (
    nombre, 
    descripcion, 
    horario, 
    duracion, 
    cuposMaximos, 
    precioBase, 
    fechaCreacion, 
    fechaActualizacion
) VALUES (?, ?, ?, ?, ?, ?, now(), now());

-- name: GetAllTours :many
SELECT * FROM Tour;

-- name: GetTourById :one
SELECT * FROM Tour WHERE idTour = ?;

-- name: UpdateTour :execresult
UPDATE Tour 
SET nombre = ?, 
    descripcion = ?, 
    horario = ?, 
    duracion = ?, 
    cuposMaximos = ?, 
    precioBase = ?, 
    fechaActualizacion = now()
WHERE idTour = ?;

-- name: DeleteTour :execresult
DELETE FROM Tour WHERE idTour = ?;
