-- name: CreateVehiculo :execresult
INSERT INTO Vehiculo (
    idChofer, 
    matricula, 
    capacidad, 
    modelo, 
    fechaCreacion, 
    fechaActualizacion
) VALUES (?, ?, ?, ?, now(), now());

-- name: GetAllVehiculos :many
SELECT 
    v.idVehiculo, 
    v.matricula, 
    v.capacidad, 
    v.modelo, 
    v.idChofer, 
    c.nombre AS nombre_chofer,
    c.telefono AS telefono_chofer,
    v.fechaCreacion,
    v.fechaActualizacion
FROM Vehiculo v
JOIN Chofer c ON v.idChofer = c.idChofer;

-- name: GetVehiculosByChofer :many
SELECT * FROM Vehiculo 
WHERE idChofer = ?;

-- name: UpdateVehiculo :execresult
UPDATE Vehiculo 
SET idChofer = ?,
    matricula = ?, 
    capacidad = ?, 
    modelo = ?, 
    fechaActualizacion = now()
WHERE idVehiculo = ?;

-- name: DeleteVehiculo :execresult
DELETE FROM Vehiculo WHERE idVehiculo = ?;
