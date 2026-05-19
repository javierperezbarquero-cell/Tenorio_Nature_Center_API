-- name: GetAllDetalleReserva :many
SELECT
    dr.idDetalleReserva,
    dr.idReserva,
    dr.fechaTour,
    dr.precioUnitario,
    dr.fechaCreacion,
    dr.fechaActualizacion,

    -- Estado de la Reserva padre
    e.nombre    AS nombreEstado,

    -- Tour
    t.nombre    AS tourNombre,
    t.horario,
    t.duracion,
    t.precioBase,

    -- Guia
    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,

    -- Chofer
    ch.nombre       AS choferNombre,
    ch.telefono     AS choferTelefono,
    ch.tipoLicencia AS choferLicencia,

    -- Vehiculo del Chofer
    v.matricula     AS vehiculoMatricula,
    v.modelo        AS vehiculoModelo,
    v.capacidad     AS vehiculoCapacidad,

    -- Ubicacion
    u.nombre    AS ubicacionNombre,
    u.direccion,

    -- Idioma
    i.nombre    AS idiomaNombre

FROM DetalleReserva dr

JOIN Reserva r      ON dr.idReserva     = r.idReserva
JOIN EstadoReserva e ON r.idEstadoReserva = e.idEstadoReserva
JOIN Tour t         ON dr.idTour        = t.idTour
JOIN Guia g         ON dr.idGuia        = g.idGuia
JOIN Chofer ch      ON dr.idChofer      = ch.idChofer
JOIN Ubicacion u    ON dr.idUbicacion   = u.idUbicacion
JOIN Idioma i       ON dr.idIdioma      = i.idIdioma
LEFT JOIN Vehiculo v ON ch.idChofer     = v.idChofer;


-- name: GetDetalleReservaById :one
SELECT
    dr.idDetalleReserva,
    dr.idReserva,
    dr.fechaTour,
    dr.precioUnitario,
    dr.fechaCreacion,
    dr.fechaActualizacion,

    e.nombre    AS nombreEstado,

    t.nombre    AS tourNombre,
    t.horario,
    t.duracion,
    t.precioBase,

    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,

    ch.nombre       AS choferNombre,
    ch.telefono     AS choferTelefono,
    ch.tipoLicencia AS choferLicencia,

    v.matricula     AS vehiculoMatricula,
    v.modelo        AS vehiculoModelo,
    v.capacidad     AS vehiculoCapacidad,

    u.nombre    AS ubicacionNombre,
    u.direccion,

    i.nombre    AS idiomaNombre

FROM DetalleReserva dr

JOIN Reserva r       ON dr.idReserva      = r.idReserva
JOIN EstadoReserva e ON r.idEstadoReserva = e.idEstadoReserva
JOIN Tour t          ON dr.idTour         = t.idTour
JOIN Guia g          ON dr.idGuia         = g.idGuia
JOIN Chofer ch       ON dr.idChofer       = ch.idChofer
JOIN Ubicacion u     ON dr.idUbicacion    = u.idUbicacion
JOIN Idioma i        ON dr.idIdioma       = i.idIdioma
LEFT JOIN Vehiculo v ON ch.idChofer       = v.idChofer

WHERE dr.idDetalleReserva = ?;


-- name: GetDetalleReservaByReserva :many
SELECT
    dr.idDetalleReserva,
    dr.idReserva,
    dr.fechaTour,
    dr.precioUnitario,
    dr.fechaCreacion,
    dr.fechaActualizacion,

    e.nombre    AS nombreEstado,

    t.nombre    AS tourNombre,
    t.horario,
    t.duracion,
    t.precioBase,

    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,

    ch.nombre       AS choferNombre,
    ch.telefono     AS choferTelefono,
    ch.tipoLicencia AS choferLicencia,

    v.matricula     AS vehiculoMatricula,
    v.modelo        AS vehiculoModelo,
    v.capacidad     AS vehiculoCapacidad,

    u.nombre    AS ubicacionNombre,
    u.direccion,

    i.nombre    AS idiomaNombre

FROM DetalleReserva dr

JOIN Reserva r       ON dr.idReserva      = r.idReserva
JOIN EstadoReserva e ON r.idEstadoReserva = e.idEstadoReserva
JOIN Tour t          ON dr.idTour         = t.idTour
JOIN Guia g          ON dr.idGuia         = g.idGuia
JOIN Chofer ch       ON dr.idChofer       = ch.idChofer
JOIN Ubicacion u     ON dr.idUbicacion    = u.idUbicacion
JOIN Idioma i        ON dr.idIdioma       = i.idIdioma
LEFT JOIN Vehiculo v ON ch.idChofer       = v.idChofer

WHERE dr.idReserva = ?;


-- name: CreateDetalleReserva :execresult
INSERT INTO DetalleReserva (
    idReserva,
    idTour,
    idGuia,
    idChofer,
    idUbicacion,
    idIdioma,
    fechaTour,
    precioUnitario,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, now(), now());


-- name: UpdateDetalleReserva :execresult
UPDATE DetalleReserva
SET idTour             = ?,
    idGuia             = ?,
    idChofer           = ?,
    idUbicacion        = ?,
    idIdioma           = ?,
    fechaTour          = ?,
    precioUnitario     = ?,
    fechaActualizacion = now()
WHERE idDetalleReserva = ?;


-- name: DeleteDetalleReserva :execresult
DELETE FROM DetalleReserva WHERE idDetalleReserva = ?;