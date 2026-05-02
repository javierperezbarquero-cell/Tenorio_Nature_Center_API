-- name: CreateTransporte :execresult
INSERT INTO transporte (idChofer, fechaCreacion, fechaActualizacion)
VALUES (?, NOW(), NOW());

-- name: GetAllTransporte :many
SELECT * FROM transporte;

-- name: GetTransporteById :one
SELECT * FROM transporte WHERE idTransporte = ?;

-- name: UpdateTransporte :execresult
UPDATE transporte
SET idChofer = ?,
    fechaActualizacion = NOW()
WHERE idTransporte = ?;

-- name: DeleteTransporte :execresult
DELETE FROM transporte WHERE idTransporte = ?;