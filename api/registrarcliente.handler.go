package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registrarClienteRequest struct {
	Nombre     string `json:"nombre"     binding:"required"`
	Apellido   string `json:"apellido"   binding:"required"`
	Correo     string `json:"correo"     binding:"required,email"`
	Contrasena string `json:"contrasena" binding:"required,min=6"`

	Identificador string `json:"identificador" binding:"required"`
	FechaNac      string `json:"fechaNac"      binding:"required"`
	Telefono      string `json:"telefono"      binding:"required"`
	Nacionalidad  string `json:"nacionalidad"  binding:"required"`
}

func isDuplicateEntry(err error) bool {
	return strings.Contains(err.Error(), "Duplicate entry")
}

func (server *Server) registrarCliente(ctx *gin.Context) {

	var req registrarClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fechaNac, err := parsearFecha(req.FechaNac, "fechaNac")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var idUsuarioGenerado int64

	err = server.ExecTx(ctx, func(q *dto.Queries) error {

		resultUsuario, err := q.CreateUsuario(ctx, dto.CreateUsuarioParams{
			Nombre:     req.Nombre,
			Apellido:   sql.NullString{String: req.Apellido, Valid: true},
			Rol:        sql.NullString{String: "Cliente", Valid: true},
			Correo:     req.Correo,
			Contrasena: string(hash),
		})
		if err != nil {
			return err
		}

		idUsuario, err := resultUsuario.LastInsertId()
		if err != nil {
			return err
		}
		idUsuarioGenerado = idUsuario

		idUsuario32 := int32(idUsuario)
		_, err = q.CreateCliente(ctx, dto.CreateClienteParams{
			Idusuario:     sql.NullInt32{Int32: idUsuario32, Valid: true},
			Nombre:        req.Nombre + " " + req.Apellido,
			Identificador: req.Identificador,
			Fechanac:      fechaNac,
			Telefono:      req.Telefono,
			Nacionalidad:  req.Nacionalidad,
			Fecharegistro: time.Now(),
		})
		return err
	})

	if err != nil {
		if isDuplicateEntry(err) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "El correo o la identificación ya están registrados"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "Cliente registrado correctamente",
		"idUsuario": idUsuarioGenerado,
	})
}