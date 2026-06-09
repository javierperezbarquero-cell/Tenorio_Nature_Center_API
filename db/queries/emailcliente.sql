-- name: GetAllEmailCliente :many
SELECT
    ec.idEmailCliente,
    ec.idCliente,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion,

    -- Cliente
    c.nombre        AS clienteNombre,
    c.identificador AS clienteIdentificador

FROM EmailCliente ec

JOIN Cliente c ON ec.idCliente = c.idCliente;


-- name: GetEmailClienteById :one
SELECT
    ec.idEmailCliente,
    ec.idCliente,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion,

    c.nombre        AS clienteNombre,
    c.identificador AS clienteIdentificador

FROM EmailCliente ec

JOIN Cliente c ON ec.idCliente = c.idCliente

WHERE ec.idEmailCliente = ?;


-- name: GetEmailClienteByCliente :many
SELECT
    ec.idEmailCliente,
    ec.idCliente,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion

FROM EmailCliente ec

WHERE ec.idCliente = ?;


-- name: CreateEmailCliente :execresult
INSERT INTO EmailCliente (
    idCliente,
    email,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, now(), now());


-- name: UpdateEmailCliente :execresult
UPDATE EmailCliente
SET email              = ?,
    fechaActualizacion = now()
WHERE idEmailCliente = ?;


-- name: DeleteEmailCliente :execresult
DELETE FROM EmailCliente WHERE idEmailCliente = ?;