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
 
    -- Reserva
    r.cantidadPersonas,
 
    -- Cliente
    c.nombre AS clienteNombre,
    c.telefono AS clienteTelefono,
 
    -- Tour
    t.nombre AS tourNombre,
 
    -- Estado Pago
    e.nombre AS nombreEstado
 
FROM Factura f
 
JOIN Reserva r ON f.idReserva = r.idReserva
JOIN Cliente c ON r.idCliente = c.idCliente
JOIN Tour t ON r.idTour = t.idTour
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