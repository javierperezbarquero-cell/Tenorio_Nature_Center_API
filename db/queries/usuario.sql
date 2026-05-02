-- name: GetUserByEmail :one
SELECT * FROM usuarios WHERE correo=? limit 1;

-- name: CreateUsuario :execresult
INSERT INTO usuarios (nombre, apellido, rol, correo, contrasena, descripcion, imagen)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateUsuario :exec
UPDATE usuarios SET 
    nombre = ?,
    apellido = ?,
    rol = ?,
    correo = ?,
    contrasena = ?,
    descripcion = ?,
    imagen = ?
WHERE idusuario = ?;

-- name: DeleteUsuario :exec
DELETE FROM usuarios WHERE idusuario = ?;