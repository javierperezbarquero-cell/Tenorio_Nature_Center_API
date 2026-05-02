-- name: GetUserByEmail :one
SELECT * FROM usuarios WHERE correo=? limit 1;

-- name: CreateUsuario :execresult
INSERT INTO usuarios (nombre, apellido, rol, correo, contrasena, descripcion, imagen, fechacreacion, fechaactualizacion)
VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateUsuario :exec
UPDATE usuarios SET 
    nombre = ?,
    apellido = ?,
    rol = ?,
    correo = ?,
    contrasena = ?,
    descripcion = ?,
    imagen = ?,
    fechaactualizacion = NOW()
WHERE idusuario = ?;

-- name: DeleteUsuario :exec
DELETE FROM usuarios WHERE idusuario = ?;