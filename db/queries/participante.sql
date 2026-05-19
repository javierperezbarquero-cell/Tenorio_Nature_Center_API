-- name: GetAllParticipante :many
SELECT
    p.idParticipante,
    p.idCliente,
    p.idReserva,
    p.fechaCreacion,
    p.fechaActualizacion,

    -- Cliente
    c.nombre            AS clienteNombre,
    c.identificador     AS clienteIdentificador,
    c.telefono          AS clienteTelefono,
    c.nacionalidad      AS clienteNacionalidad,
    c.fechaNac          AS clienteFechaNac,

    -- Empresa del cliente (si tiene)
    ec.razonSocial      AS empresaNombre,
    ec.cedulaJuridica   AS empresaCedula,

    -- Reserva
    e.nombre            AS nombreEstado,

    -- Tour asociado a la reserva
    t.nombre            AS tourNombre,
    dr.fechaTour

FROM Participante p

JOIN Cliente c              ON p.idCliente          = c.idCliente
JOIN Reserva r              ON p.idReserva          = r.idReserva
JOIN EstadoReserva e        ON r.idEstadoReserva    = e.idEstadoReserva
JOIN DetalleReserva dr      ON r.idReserva          = dr.idReserva
JOIN Tour t                 ON dr.idTour            = t.idTour
LEFT JOIN EmpresaCliente ec ON c.idEmpresaCliente   = ec.idEmpresaCliente;


-- name: GetParticipanteById :one
SELECT
    p.idParticipante,
    p.idCliente,
    p.idReserva,
    p.fechaCreacion,
    p.fechaActualizacion,

    c.nombre            AS clienteNombre,
    c.identificador     AS clienteIdentificador,
    c.telefono          AS clienteTelefono,
    c.nacionalidad      AS clienteNacionalidad,
    c.fechaNac          AS clienteFechaNac,

    ec.razonSocial      AS empresaNombre,
    ec.cedulaJuridica   AS empresaCedula,

    e.nombre            AS nombreEstado,

    t.nombre            AS tourNombre,
    dr.fechaTour

FROM Participante p

JOIN Cliente c              ON p.idCliente          = c.idCliente
JOIN Reserva r              ON p.idReserva          = r.idReserva
JOIN EstadoReserva e        ON r.idEstadoReserva    = e.idEstadoReserva
JOIN DetalleReserva dr      ON r.idReserva          = dr.idReserva
JOIN Tour t                 ON dr.idTour            = t.idTour
LEFT JOIN EmpresaCliente ec ON c.idEmpresaCliente   = ec.idEmpresaCliente

WHERE p.idParticipante = ?;


-- name: GetParticipanteByReserva :many
SELECT
    p.idParticipante,
    p.idCliente,
    p.idReserva,
    p.fechaCreacion,
    p.fechaActualizacion,

    c.nombre            AS clienteNombre,
    c.identificador     AS clienteIdentificador,
    c.telefono          AS clienteTelefono,
    c.nacionalidad      AS clienteNacionalidad,
    c.fechaNac          AS clienteFechaNac,

    ec.razonSocial      AS empresaNombre,
    ec.cedulaJuridica   AS empresaCedula

FROM Participante p

JOIN Cliente c              ON p.idCliente        = c.idCliente
LEFT JOIN EmpresaCliente ec ON c.idEmpresaCliente = ec.idEmpresaCliente

WHERE p.idReserva = ?;


-- name: GetParticipanteByCliente :many
SELECT
    p.idParticipante,
    p.idCliente,
    p.idReserva,
    p.fechaCreacion,
    p.fechaActualizacion,

    e.nombre    AS nombreEstado,
    t.nombre    AS tourNombre,
    dr.fechaTour,
    dr.precioUnitario

FROM Participante p

JOIN Reserva r          ON p.idReserva          = r.idReserva
JOIN EstadoReserva e    ON r.idEstadoReserva    = e.idEstadoReserva
JOIN DetalleReserva dr  ON r.idReserva          = dr.idReserva
JOIN Tour t             ON dr.idTour            = t.idTour

WHERE p.idCliente = ?;


-- name: CreateParticipante :execresult
INSERT INTO Participante (
    idCliente,
    idReserva,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, now(), now());


-- name: DeleteParticipante :execresult
DELETE FROM Participante WHERE idParticipante = ?;