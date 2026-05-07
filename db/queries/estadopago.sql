-- name: GetAllEstadoPago :many
SELECT * FROM EstadoPago;
 
-- name: GetEstadoPagoById :one
SELECT * FROM EstadoPago WHERE idEstadoPago = ?;
 
-- name: CreateEstadoPago :execresult
INSERT INTO EstadoPago (nombre, descripcion, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());
 
-- name: UpdateEstadoPago :execresult
UPDATE EstadoPago
SET nombre = ?,
    descripcion = ?,
    fechaActualizacion = now()
WHERE idEstadoPago = ?;
 
-- name: DeleteEstadoPago :execresult
DELETE FROM EstadoPago WHERE idEstadoPago = ?;