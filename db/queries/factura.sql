-- name: GetAllFactura :many
SELECT
    f.idFactura,
    f.numeroFactura,
    f.fechaFactura,
    f.metodoPago,
    f.moneda,
    f.fechaPago,
    f.subtotal,
    f.impuesto,
    f.descuento,
    f.precioTotal,
    f.fechaCreacion,
    f.fechaActualizacion,

    -- Cantidad de participantes cubiertos por esta factura
    COUNT(fp.idParticipante) AS cantidadPersonas,

    -- Nombres de los clientes cubiertos (pueden ser varios)
    GROUP_CONCAT(c.nombre ORDER BY c.nombre SEPARATOR ', ') AS clientesNombre,

    -- Telefonos de los clientes cubiertos
    GROUP_CONCAT(c.telefono ORDER BY c.nombre SEPARATOR ', ') AS clientesTelefono,

    -- Tour se obtiene el idReserva desde Participante
    (SELECT MAX(t.nombre)
    FROM DetalleReserva dr
    JOIN Tour t ON t.idTour = dr.idTour
    WHERE dr.idReserva = p.idReserva) AS tourNombre,
    -- Estado de pago
    e.nombre AS nombreEstado

FROM Factura f
JOIN FacturaParticipante fp ON fp.idFactura     = f.idFactura
JOIN Participante        p  ON p.idParticipante  = fp.idParticipante
JOIN Cliente             c  ON c.idCliente       = p.idCliente
JOIN EstadoPago          e  ON e.idEstadoPago    = f.idEstadoPago
GROUP BY
    f.idFactura,
    f.numeroFactura,
    f.fechaFactura,
    f.metodoPago,
    f.moneda,
    f.fechaPago,
    f.subtotal,
    f.impuesto,
    f.descuento,
    f.precioTotal,
    f.fechaCreacion,
    f.fechaActualizacion,
    e.nombre;


-- name: GetFacturaById :one
SELECT
    f.idFactura,
    f.numeroFactura,
    f.fechaFactura,
    f.metodoPago,
    f.moneda,
    f.fechaPago,
    f.subtotal,
    f.impuesto,
    f.descuento,
    f.precioTotal,
    f.fechaCreacion,
    f.fechaActualizacion,

    COUNT(fp.idParticipante) AS cantidadPersonas,
    GROUP_CONCAT(c.nombre ORDER BY c.nombre SEPARATOR ', ') AS clientesNombre,
    GROUP_CONCAT(c.telefono ORDER BY c.nombre SEPARATOR ', ') AS clientesTelefono,

    (SELECT MAX(t.nombre)
    FROM DetalleReserva dr
    JOIN Tour t ON t.idTour = dr.idTour
    WHERE dr.idReserva = p.idReserva) AS tourNombre,
    e.nombre AS nombreEstado

FROM Factura f
JOIN FacturaParticipante fp ON fp.idFactura     = f.idFactura
JOIN Participante        p  ON p.idParticipante  = fp.idParticipante
JOIN Cliente             c  ON c.idCliente       = p.idCliente
JOIN EstadoPago          e  ON e.idEstadoPago    = f.idEstadoPago
WHERE f.idFactura = ?
GROUP BY
    f.idFactura,
    f.numeroFactura,
    f.fechaFactura,
    f.metodoPago,
    f.moneda,
    f.fechaPago,
    f.subtotal,
    f.impuesto,
    f.descuento,
    f.precioTotal,
    f.fechaCreacion,
    f.fechaActualizacion,
    e.nombre;

-- name: GetReservasDisponiblesParaFacturar :many
SELECT
    r.idReserva,
    r.idEstadoReserva,
    r.fechaCreacion,
    r.fechaActualizacion,
    GROUP_CONCAT(DISTINCT c.nombre ORDER BY c.nombre SEPARATOR ', ') AS clientenombres
FROM Reserva r
JOIN Participante p ON p.idReserva = r.idReserva
JOIN Cliente c ON c.idCliente = p.idCliente
LEFT JOIN FacturaParticipante fp ON fp.idParticipante = p.idParticipante
GROUP BY
    r.idReserva,
    r.idEstadoReserva,
    r.fechaCreacion,
    r.fechaActualizacion
HAVING SUM(CASE WHEN fp.idParticipante IS NULL THEN 1 ELSE 0 END) > 0;

-- name: GetParticiapntesSinFactura :many
SELECT p.idParticipante, c.nombre
FROM Participante p
JOIN Cliente c ON p.idCliente = c.idCliente
LEFT JOIN FacturaParticipante fp ON p.idParticipante = fp.idParticipante
WHERE p.idReserva = ?
AND fp.idParticipante IS NULL;

-- name: CreateFactura :execresult
INSERT INTO Factura (
    idEstadoPago,
    numeroFactura,
    fechaFactura,
    metodoPago,
    moneda,
    fechaPago,
    subtotal,
    impuesto,
    descuento,
    precioTotal,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, now(), now());


-- name: UpdateFactura :execresult
UPDATE Factura
SET idEstadoPago       = ?,
    numeroFactura      = ?,
    fechaFactura       = ?,
    metodoPago         = ?,
    moneda             = ?,
    fechaPago          = ?,
    subtotal           = ?,
    impuesto           = ?,
    descuento          = ?,
    precioTotal        = ?,
    fechaActualizacion = now()
WHERE idFactura = ?;


-- name: DeleteFactura :execresult
DELETE FROM Factura WHERE idFactura = ?;