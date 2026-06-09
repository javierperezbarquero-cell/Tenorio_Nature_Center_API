package api

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"rest/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	IDUsuario   int32  `json:"idUsuario"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Rol         string `json:"rol"`
	Correo      string `json:"correo"`
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
		fmt.Println("ERROR LOGIN:", err)
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
		time.Minute*15,
	)
	if err != nil {
		fmt.Println("ERROR LOGIN:", err)
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
			Correo:      user.Correo,
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
		fmt.Println("ERROR LOGIN:", err)
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
		Nombre:      req.Nombre,
		Apellido:    sql.NullString{String: req.Apellido, Valid: req.Apellido != ""},
		Rol:         sql.NullString{String: req.Rol, Valid: req.Rol != ""},
		Correo:      req.Correo,
		Contrasena:  req.Contrasena,
		Descripcion: sql.NullString{String: req.Descripcion, Valid: req.Descripcion != ""},
		Imagen:      sql.NullString{String: req.Imagen, Valid: req.Imagen != ""},
		Idusuario:   req.Idusuario,
	}

	err := server.dbtx.UpdateUsuario(ctx, arg)
	if err != nil {
		fmt.Println("ERROR ACTUALIZAR USUARIO:", err)
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
		fmt.Println("ERROR LOGIN:", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}

func (server *Server) uploadUserImg(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file0")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	defer file.Close()

	upDir := "utils/images/users"
	if _, err := os.Stat(upDir); os.IsNotExist(err) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	filename := uuid.New().String() + "_" + filepath.Base(fileHeader.Filename)
	destinationFile, err := os.Create(filepath.Join(upDir, filename))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"filename": filename,
		"message":  "Imagen cargada exitosamente",
	})
}

type userImageRequest struct {
	Filename string `uri:"filename" binding:"required"`
}

func (server *Server) downloadUserImg(ctx *gin.Context) {
	var req userImageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fileUrl := "utils/images/users/" + req.Filename
	ctx.File(fileUrl)
}

func (server *Server) getAllUsuarios(ctx *gin.Context) {
    usuarios, err := server.dbtx.GetAllUsuarios(ctx)
    if err != nil {
        fmt.Println("ERROR GET ALL USUARIOS:", err)
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }
    ctx.JSON(http.StatusOK, usuarios)
}
