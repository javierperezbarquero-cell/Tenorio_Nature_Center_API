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

    -- Cantidad de participantes de la reserva
    (SELECT COUNT(*) 
     FROM Participante p2 
     WHERE p2.idReserva = p.idReserva) AS cantidadPersonas,

    -- Cliente dueño de esta factura
    c.nombre AS clienteNombre,
    c.telefono AS clienteTelefono,

    -- Tour: primer detalle de la reserva
    (SELECT t.nombre 
     FROM DetalleReserva dr 
     JOIN Tour t ON t.idTour = dr.idTour 
     WHERE dr.idReserva = p.idReserva 
     LIMIT 1) AS tourNombre,

    -- Estado Pago
    e.nombre AS nombreEstado

FROM Factura f
JOIN Participante p ON f.idParticipante = p.idParticipante
JOIN Cliente c ON p.idCliente = c.idCliente
JOIN EstadoPago e ON f.idEstadoPago = e.idEstadoPago;

-- name: GetFacturaById :one
SELECT * FROM Factura WHERE idFactura = ?;

-- name: CreateFactura :execresult
INSERT INTO Factura (
    idParticipante,
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
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, now(), now());

-- name: UpdateFactura :execresult
UPDATE Factura
SET idParticipante = ?,
    idEstadoPago = ?,
    numeroFactura = ?,
    fechaFactura = ?,
    metodoPago = ?,
    moneda = ?,
    fechaPago = ?,
    subtotal = ?,
    impuesto = ?,
    descuento = ?,
    precioTotal = ?,
    fechaActualizacion = now()
WHERE idFactura = ?;

-- name: DeleteFactura :execresult
DELETE FROM Factura WHERE idFactura = ?;