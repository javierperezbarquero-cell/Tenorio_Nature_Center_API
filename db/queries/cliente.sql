-- name: GetAllCliente :many
SELECT * FROM Cliente;

-- name: GetClienteById :one
SELECT * FROM Cliente WHERE idCliente = ?;

-- name: CreateCliente :execresult
INSERT INTO Cliente (idEmpresaCliente, idUsuario, nombre, identificador, fechaNac, telefono, nacionalidad, fechaRegistro, fechaCreacion, fechaActualizacion)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, now(), now());

-- name: UpdateCliente :execresult
UPDATE Cliente 
SET idEmpresaCliente = ?,
    idUsuario = ?,
    nombre = ?,
    identificador = ?,
    fechaNac = ?,
    telefono = ?, 
    nacionalidad = ?,
    fechaRegistro = ?, 
    fechaActualizacion = now()
WHERE idCliente = ?;

-- name: DeleteCliente :execresult
DELETE FROM Cliente WHERE idCliente = ?;

-- name: GetClienteByUsuario :one
SELECT * FROM Cliente
WHERE idUsuario = ?;