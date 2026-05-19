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

    -- Cantidad de participantes calculada desde Participante
    (SELECT COUNT(*) 
     FROM Participante p 
     WHERE p.idReserva = r.idReserva) AS cantidadPersonas,

    -- Cliente: primer participante de la reserva
    (SELECT c.nombre 
     FROM Participante p 
     JOIN Cliente c ON c.idCliente = p.idCliente 
     WHERE p.idReserva = r.idReserva 
     LIMIT 1) AS clienteNombre,

    (SELECT c.telefono 
     FROM Participante p 
     JOIN Cliente c ON c.idCliente = p.idCliente 
     WHERE p.idReserva = r.idReserva 
     LIMIT 1) AS clienteTelefono,

    -- Tour: primer detalle de la reserva
    (SELECT t.nombre 
     FROM DetalleReserva dr 
     JOIN Tour t ON t.idTour = dr.idTour 
     WHERE dr.idReserva = r.idReserva 
     LIMIT 1) AS tourNombre,

    -- Estado Pago
    e.nombre AS nombreEstado

FROM Factura f
JOIN Reserva r ON f.idReserva = r.idReserva
JOIN EstadoPago e ON f.idEstadoPago = e.idEstadoPago;

-- name: GetFacturaById :one
SELECT * FROM Factura WHERE idFactura = ?;

-- name: CreateFactura :execresult
INSERT INTO Factura (
    idReserva,
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
SET idReserva = ?,
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