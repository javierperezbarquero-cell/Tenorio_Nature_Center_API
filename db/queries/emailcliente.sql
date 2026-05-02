-- name: GetAllEmailCliente :many
SELECT * FROM emailcliente;

-- name: GetEmailClienteById :one
SELECT * FROM emailcliente WHERE idEmailCliente = ?;

-- name: CreateEmailCliente :execresult
INSERT INTO emailcliente (email, idCliente, fechaCreacion, fechaActualizacion)
VALUES (?, ?, now(), now());

-- name: UpdateEmailCliente :execresult
UPDATE emailcliente 
SET email = ?, 
    idCliente = ?,
    fechaActualizacion = now()
WHERE idEmailCliente = ?;

-- name: DeleteEmailCliente :execresult
DELETE FROM emailcliente WHERE idEmailCliente = ?;