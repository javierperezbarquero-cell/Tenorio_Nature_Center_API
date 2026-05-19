-- name: GetAllReserva :many
SELECT 
    r.idReserva,
    COUNT(p.idParticipante)         AS cantidadPersonas,
    dr.fechaTour,
    r.fechaCreacion,
    r.fechaActualizacion,

    -- Clientes (múltiples por reserva)
    GROUP_CONCAT(DISTINCT c.nombre SEPARATOR ', ')   AS clienteNombres,
    GROUP_CONCAT(DISTINCT c.telefono SEPARATOR ', ')  AS clienteTelefonos,

    -- Tour
    t.nombre    AS tourNombre,
    t.horario,
    t.duracion,

    -- Guia
    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,

    -- Chofer (directo desde DetalleReserva)
    ch.nombre   AS choferNombre,
    ch.telefono AS choferTelefono,

    -- Ubicacion
    u.nombre    AS ubicacionNombre,
    u.direccion,

    -- Idioma
    i.nombre    AS idiomaNombre,

    -- Estado Reserva
    e.nombre    AS nombreEstado

FROM Reserva r

JOIN DetalleReserva dr  ON r.idReserva       = dr.idReserva
JOIN Tour t             ON dr.idTour         = t.idTour
JOIN Guia g             ON dr.idGuia         = g.idGuia
JOIN Chofer ch          ON dr.idChofer       = ch.idChofer
JOIN Ubicacion u        ON dr.idUbicacion    = u.idUbicacion
JOIN Idioma i           ON dr.idIdioma       = i.idIdioma
JOIN EstadoReserva e    ON r.idEstadoReserva = e.idEstadoReserva
LEFT JOIN Participante p ON r.idReserva      = p.idReserva
LEFT JOIN Cliente c      ON p.idCliente      = c.idCliente

GROUP BY
    r.idReserva,
    dr.idDetalleReserva,
    dr.fechaTour,
    r.fechaCreacion,
    r.fechaActualizacion,
    t.nombre, t.horario, t.duracion,
    g.nombre, g.telefono,
    ch.nombre, ch.telefono,
    u.nombre, u.direccion,
    i.nombre,
    e.nombre;


-- name: GetReservaById :one
SELECT 
    r.idReserva,
    COUNT(p.idParticipante)         AS cantidadPersonas,
    dr.fechaTour,
    r.fechaCreacion,
    r.fechaActualizacion,

    GROUP_CONCAT(DISTINCT c.nombre SEPARATOR ', ')   AS clienteNombres,
    GROUP_CONCAT(DISTINCT c.telefono SEPARATOR ', ')  AS clienteTelefonos,

    t.nombre    AS tourNombre,
    t.horario,
    t.duracion,

    g.nombre    AS guiaNombre,
    g.telefono  AS guiaTelefono,

    ch.nombre   AS choferNombre,
    ch.telefono AS choferTelefono,

    u.nombre    AS ubicacionNombre,
    u.direccion,

    i.nombre    AS idiomaNombre,
    e.nombre    AS nombreEstado

FROM Reserva r

JOIN DetalleReserva dr  ON r.idReserva       = dr.idReserva
JOIN Tour t             ON dr.idTour         = t.idTour
JOIN Guia g             ON dr.idGuia         = g.idGuia
JOIN Chofer ch          ON dr.idChofer       = ch.idChofer
JOIN Ubicacion u        ON dr.idUbicacion    = u.idUbicacion
JOIN Idioma i           ON dr.idIdioma       = i.idIdioma
JOIN EstadoReserva e    ON r.idEstadoReserva = e.idEstadoReserva
LEFT JOIN Participante p ON r.idReserva      = p.idReserva
LEFT JOIN Cliente c      ON p.idCliente      = c.idCliente

WHERE r.idReserva = ?

GROUP BY
    r.idReserva,
    dr.idDetalleReserva,
    dr.fechaTour,
    r.fechaCreacion,
    r.fechaActualizacion,
    t.nombre, t.horario, t.duracion,
    g.nombre, g.telefono,
    ch.nombre, ch.telefono,
    u.nombre, u.direccion,
    i.nombre,
    e.nombre;


-- name: CreateReserva :execresult
INSERT INTO Reserva (
    idEstadoReserva,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, now(), now());


-- name: UpdateReserva :execresult
UPDATE Reserva
SET idEstadoReserva    = ?,
    fechaActualizacion = now()
WHERE idReserva = ?;


-- name: DeleteReserva :execresult
DELETE FROM Reserva WHERE idReserva = ?;