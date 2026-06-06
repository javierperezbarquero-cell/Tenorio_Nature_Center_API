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
    (SELECT t.nombre
     FROM DetalleReserva dr
     JOIN Tour t ON t.idTour = dr.idTour
     WHERE dr.idReserva = p.idReserva
     LIMIT 1) AS tourNombre,

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

    (SELECT t.nombre
     FROM DetalleReserva dr
     JOIN Tour t ON t.idTour = dr.idTour
     WHERE dr.idReserva = p.idReserva
     LIMIT 1) AS tourNombre,

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