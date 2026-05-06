package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createClienteRequest struct {
	Nombre        string    `json:"nombre" binding:"required"`
	Telefono      string     `json:"telefono" binding:"required"`
	Nacionalidad  string    `json:"nacionalidad" binding:"required"`
	FechaRegistro string `json:"fechaRegistro" binding:"required"`
}

type updateClienteRequest struct {
	IdCliente     int32     `json:"idCliente" binding:"required"`
	Nombre        string    `json:"nombre" binding:"required"`
	Telefono      string     `json:"telefono" binding:"required"`
	Nacionalidad  string    `json:"nacionalidad" binding:"required"`
	FechaRegistro string `json:"fechaRegistro" binding:"required"`
}

func (server *Server) createCliente(ctx *gin.Context) {
	var req createClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fechaReg, err := parsearFecha(req.FechaRegistro, "fechaNac")
	if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}
	args := dto.CreateClienteParams{
		Nombre:        req.Nombre,
		Telefono:      req.Telefono,
		Nacionalidad:  req.Nacionalidad,
		Fecharegistro: fechaReg,
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
	fechaReg, err := parsearFecha(req.FechaRegistro, "fechaNac")
	if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}

	args := dto.UpdateClienteParams{
		Nombre:        req.Nombre,
		Telefono:      req.Telefono,
		Nacionalidad:  req.Nacionalidad,
		Fecharegistro: fechaReg,
		Idcliente:     req.IdCliente,
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
