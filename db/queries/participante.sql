-- name: GetAllParticipante :many
SELECT * FROM Participante;

-- name: GetParticipanteById :one
SELECT * FROM Participante WHERE idParticipante = ?;

-- name: CreateParticipante :execresult
INSERT INTO Participante (
    idReserva,
    nombre,
    fechaNac,
    nacionalidad,
    telefono,
    fechaCreacion,
    fechaActualizacion
)
VALUES (?, ?, ?, ?, ?, now(), now());

-- name: UpdateParticipante :execresult
UPDATE Participante
SET idReserva = ?,
    nombre = ?,
    fechaNac = ?,
    nacionalidad = ?,
    telefono = ?,
    fechaActualizacion = now()
WHERE idParticipante = ?;

-- name: DeleteParticipante :execresult
DELETE FROM Participante WHERE idParticipante = ?;