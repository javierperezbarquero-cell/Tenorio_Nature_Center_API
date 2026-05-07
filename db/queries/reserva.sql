-- name: GetAllReserva :many
SELECT 
    r.idReserva,
    r.cantidadPersonas,
    r.fechaTour,
    r.fechaCreacion,
    r.fechaActualizacion,
 
    -- Cliente
    c.nombre AS clienteNombre,
    c.telefono AS clienteTelefono,
 
    -- Tour
    t.nombre AS tourNombre,
    t.horario,
    t.duracion,
 
    -- Guia
    g.nombre AS guiaNombre,
    g.telefono AS guiaTelefono,
 
    -- Chofer (desde Transporte)
    ch.nombre AS choferNombre,
    ch.telefono AS choferTelefono,
 
    -- Ubicacion
    u.nombre AS ubicacionNombre,
    u.direccion,
 
    -- Idioma
    i.nombre AS idiomaNombre,
 
    -- Estado Reserva
    e.nombre AS nombreEstado
 
FROM Reserva r
 
JOIN Cliente c ON r.idCliente = c.idCliente
JOIN Tour t ON r.idTour = t.idTour
JOIN Guia g ON r.idGuia = g.idGuia
JOIN Transporte tr ON r.idTransporte = tr.idTransporte
JOIN Chofer ch ON tr.idChofer = ch.idChofer
JOIN Ubicacion u ON r.idUbicacion = u.idUbicacion
JOIN Idioma i ON r.idIdioma = i.idIdioma
JOIN EstadoReserva e ON r.idEstadoReserva = e.idEstadoReserva;
 
-- name: GetReservaById :one
SELECT * FROM Reserva WHERE idReserva = ?;
 
-- name: CreateReserva :execresult
INSERT INTO Reserva (
    idCliente,
    idTour,
    idGuia,
    idTransporte,
    idUbicacion,
    idIdioma,
    idEstadoReserva,
    cantidadPersonas,
    fechaTour,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, now(), now());
 
-- name: UpdateReserva :execresult
UPDATE Reserva
SET idCliente = ?,
    idTour = ?,
    idGuia = ?,
    idTransporte = ?,
    idUbicacion = ?,
    idIdioma = ?,
    idEstadoReserva = ?,
    cantidadPersonas = ?,
    fechaTour = ?,
    fechaActualizacion = now()
WHERE idReserva = ?;
 
-- name: DeleteReserva :execresult
DELETE FROM Reserva WHERE idReserva = ?;