-- name: GetParticipantesByFactura :many
SELECT
    fp.idFacturaParticipante,
    fp.idFactura,
    fp.idParticipante,
    c.nombre      AS clienteNombre,
    c.telefono    AS clienteTelefono,
    c.identificador AS clienteIdentificador,
    fp.fechaCreacion,
    fp.fechaActualizacion
FROM FacturaParticipante fp
JOIN Participante p ON p.idParticipante = fp.idParticipante
JOIN Cliente      c ON c.idCliente      = p.idCliente
WHERE fp.idFactura = ?;
 
 
-- name: GetFacturaByParticipante :one
SELECT
    fp.idFacturaParticipante,
    fp.idFactura,
    fp.idParticipante,
    fp.fechaCreacion,
    fp.fechaActualizacion
FROM FacturaParticipante fp
WHERE fp.idParticipante = ?;
 
 
-- name: CreateFacturaParticipante :execresult
INSERT INTO FacturaParticipante (
    idFactura,
    idParticipante,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, now(), now());
 
 
-- name: UpdateFacturaParticipante :execresult
UPDATE FacturaParticipante
SET idFactura          = ?,
    fechaActualizacion = now()
WHERE idParticipante = ?;
 
 
-- name: DeleteFacturaParticipante :execresult
DELETE FROM FacturaParticipante
WHERE idFacturaParticipante = ?;
 
 
-- name: DeleteFacturaParticipanteByFactura :execresult
DELETE FROM FacturaParticipante
WHERE idFactura = ?;