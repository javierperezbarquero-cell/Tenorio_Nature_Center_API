-- name: GetAllEmpresaCliente :many
SELECT
    ec.idEmpresaCliente,
    ec.razonSocial,
    ec.cedulaJuridica,
    ec.telefono,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion,

    -- Cantidad de clientes asociados a la empresa
    COUNT(c.idCliente) AS totalClientes

FROM EmpresaCliente ec

LEFT JOIN Cliente c ON ec.idEmpresaCliente = c.idEmpresaCliente

GROUP BY
    ec.idEmpresaCliente,
    ec.razonSocial,
    ec.cedulaJuridica,
    ec.telefono,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion;


-- name: GetEmpresaClienteById :one
SELECT
    ec.idEmpresaCliente,
    ec.razonSocial,
    ec.cedulaJuridica,
    ec.telefono,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion,

    COUNT(c.idCliente) AS totalClientes

FROM EmpresaCliente ec

LEFT JOIN Cliente c ON ec.idEmpresaCliente = c.idEmpresaCliente

WHERE ec.idEmpresaCliente = ?

GROUP BY
    ec.idEmpresaCliente,
    ec.razonSocial,
    ec.cedulaJuridica,
    ec.telefono,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion;


-- name: GetEmpresaClienteByCedula :one
SELECT
    ec.idEmpresaCliente,
    ec.razonSocial,
    ec.cedulaJuridica,
    ec.telefono,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion,

    COUNT(c.idCliente) AS totalClientes

FROM EmpresaCliente ec

LEFT JOIN Cliente c ON ec.idEmpresaCliente = c.idEmpresaCliente

WHERE ec.cedulaJuridica = ?

GROUP BY
    ec.idEmpresaCliente,
    ec.razonSocial,
    ec.cedulaJuridica,
    ec.telefono,
    ec.email,
    ec.fechaCreacion,
    ec.fechaActualizacion;


-- name: CreateEmpresaCliente :execresult
INSERT INTO EmpresaCliente (
    razonSocial,
    cedulaJuridica,
    telefono,
    email,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, ?, ?, now(), now());


-- name: UpdateEmpresaCliente :execresult
UPDATE EmpresaCliente
SET razonSocial        = ?,
    cedulaJuridica     = ?,
    telefono           = ?,
    email              = ?,
    fechaActualizacion = now()
WHERE idEmpresaCliente = ?;


-- name: DeleteEmpresaCliente :execresult
DELETE FROM EmpresaCliente WHERE idEmpresaCliente = ?;