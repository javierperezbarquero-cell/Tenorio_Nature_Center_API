package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createClienteRequest struct {
	IdEmpresaCliente *int32 `json:"idEmpresaCliente"`
	IdUsuario        *int32 `json:"idUsuario"`
	Nombre           string `json:"nombre"        binding:"required"`
	Identificador    string `json:"identificador" binding:"required"`
	FechaNac         string `json:"fechaNac"      binding:"required"`
	Telefono         string `json:"telefono"      binding:"required"`
	Nacionalidad     string `json:"nacionalidad"  binding:"required"`
	FechaRegistro    string `json:"fechaRegistro" binding:"required"`
}

type updateClienteRequest struct {
	IdCliente        int32  `json:"idCliente" binding:"required"`
	IdEmpresaCliente *int32 `json:"idEmpresaCliente"`
	IdUsuario        *int32 `json:"idUsuario"`
	Nombre           string `json:"nombre"        binding:"required"`
	Identificador    string `json:"identificador" binding:"required"`
	FechaNac         string `json:"fechaNac"      binding:"required"`
	Telefono         string `json:"telefono"      binding:"required"`
	Nacionalidad     string `json:"nacionalidad"  binding:"required"`
	FechaRegistro    string `json:"fechaRegistro" binding:"required"`
}

func (server *Server) createCliente(ctx *gin.Context) {
	var req createClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fechaNac, err := parsearFecha(req.FechaNac, "fechaNac")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaRegistro, err := parsearFecha(req.FechaRegistro, "fechaRegistro")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.CreateClienteParams{
		Idempresacliente: toNullInt32(req.IdEmpresaCliente),
		Idusuario:        toNullInt32(req.IdUsuario),
		Nombre:           req.Nombre,
		Identificador:    req.Identificador,
		Fechanac:         fechaNac,
		Telefono:         req.Telefono,
		Nacionalidad:     req.Nacionalidad,
		Fecharegistro:    fechaRegistro,
	}

	cliente, err := server.dbtx.CreateCliente(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := cliente.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllClientes(ctx *gin.Context) {
	clientes, err := server.dbtx.GetAllCliente(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clientes)
}

func (server *Server) getClienteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	cliente, err := server.dbtx.GetClienteById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cliente)
}

func (server *Server) updateCliente(ctx *gin.Context) {
	var req updateClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fechaNac, err := parsearFecha(req.FechaNac, "fechaNac")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaRegistro, err := parsearFecha(req.FechaRegistro, "fechaRegistro")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.UpdateClienteParams{
		Idempresacliente: toNullInt32(req.IdEmpresaCliente),
		Idusuario:        toNullInt32(req.IdUsuario),
		Nombre:           req.Nombre,
		Identificador:    req.Identificador,
		Fechanac:         fechaNac,
		Telefono:         req.Telefono,
		Nacionalidad:     req.Nacionalidad,
		Fecharegistro:    fechaRegistro,
		Idcliente:        req.IdCliente,
	}

	_, err = server.dbtx.UpdateCliente(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente actualizado correctamente"})
}

func (server *Server) deleteCliente(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteCliente(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado"})
}

func (server *Server) getClienteByUsuario(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": "ID inválido"})
		return
	}

	cliente, err := server.dbtx.GetClienteByUsuario(
		ctx,
		sql.NullInt32{
			Int32: int32(id),
			Valid: true,
		},
	)

	if err != nil {

		if err == sql.ErrNoRows {

			ctx.JSON(
				http.StatusNotFound,
				gin.H{"error": "Cliente no encontrado"},
			)

			return
		}

		ctx.JSON(
			http.StatusInternalServerError,
			errorResponse(err),
		)

		return
	}

	ctx.JSON(http.StatusOK, cliente)
}
