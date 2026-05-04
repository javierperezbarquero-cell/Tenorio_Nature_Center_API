package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Correo     string `json:"correo" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

type LoginResponse struct {
	AccessToken string  `json:"access_token"`
	Payload     payload `json:"payload"`
}

type payload struct {
	IDUsuario   int32  `json:"id_usuario"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Rol         string `json:"rol"`
	Imagen      string `json:"imagen"`
	Descripcion string `json:"descripcion"`
}

type createUsuarioRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Apellido    string `json:"apellido"`
	Rol         string `json:"rol"`
	Correo      string `json:"correo" binding:"required"`
	Contrasena  string `json:"contrasena" binding:"required"`
	Descripcion string `json:"descripcion"`
	Imagen      string `json:"imagen"`
}

type updateUsuarioRequest struct {
	Idusuario   int32  `json:"idusuario" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Apellido    string `json:"apellido"`
	Rol         string `json:"rol"`
	Correo      string `json:"correo" binding:"required"`
	Contrasena  string `json:"contrasena" binding:"required"`
	Descripcion string `json:"descripcion"`
	Imagen      string `json:"imagen"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.dbtx.GetUserByEmail(ctx, req.Correo)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.Contrasena != req.Contrasena {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Contraseña incorrecta"})
		return
	}

	accessToken, err := server.tokenBuilder.CreateToken(
		user.Rol.String,
		user.Nombre,
		user.Imagen.String,
		time.Minute*5,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := LoginResponse{
		AccessToken: accessToken,
		Payload: payload{
			IDUsuario:   user.Idusuario,
			Nombre:      user.Nombre,
			Apellido:    user.Apellido.String,
			Rol:         user.Rol.String,
			Imagen:      user.Imagen.String,
			Descripcion: user.Descripcion.String,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) createUsuario(ctx *gin.Context) {
	var req createUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := dto.CreateUsuarioParams{
		Nombre:      req.Nombre,
		Apellido:    sql.NullString{String: req.Apellido, Valid: req.Apellido != ""},
		Rol:         sql.NullString{String: req.Rol, Valid: req.Rol != ""},
		Correo:      req.Correo,
		Contrasena:  req.Contrasena,
		Descripcion: sql.NullString{String: req.Descripcion, Valid: req.Descripcion != ""},
		Imagen:      sql.NullString{String: req.Imagen, Valid: req.Imagen != ""},
	}

	_, err := server.dbtx.CreateUsuario(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente"})
}

func (server *Server) updateUsuario(ctx *gin.Context) {
	var req updateUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := dto.UpdateUsuarioParams{
		Idusuario:   req.Idusuario,
		Nombre:      req.Nombre,
		Apellido:    sql.NullString{String: req.Apellido, Valid: req.Apellido != ""},
		Rol:         sql.NullString{String: req.Rol, Valid: req.Rol != ""},
		Correo:      req.Correo,
		Contrasena:  req.Contrasena,
		Descripcion: sql.NullString{String: req.Descripcion, Valid: req.Descripcion != ""},
		Imagen:      sql.NullString{String: req.Imagen, Valid: req.Imagen != ""},
	}

	err := server.dbtx.UpdateUsuario(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}

func (server *Server) deleteUsuario(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	err = server.dbtx.DeleteUsuario(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
