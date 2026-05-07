-- name: GetAllEstadoReserva :many
SELECT * FROM EstadoReserva;
 
-- name: GetEstadoReservaById :one
SELECT * FROM EstadoReserva WHERE idEstadoReserva = ?;
 
-- name: CreateEstadoReserva :execresult
INSERT INTO EstadoReserva (nombre, descripcion, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());
 
-- name: UpdateEstadoReserva :execresult
UPDATE EstadoReserva
SET nombre = ?,
    descripcion = ?,
    fechaActualizacion = now()
WHERE idEstadoReserva = ?;
 
-- name: DeleteEstadoReserva :execresult
DELETE FROM EstadoReserva WHERE idEstadoReserva = ?;