package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createEmpresaClienteRequest struct {
	RazonSocial    string `json:"razonSocial"    binding:"required"`
	CedulaJuridica string `json:"cedulaJuridica" binding:"required"`
	Telefono       string `json:"telefono"`
	Email          string `json:"email"          binding:"required"`
}

type updateEmpresaClienteRequest struct {
	IdEmpresaCliente int32  `json:"idEmpresaCliente" binding:"required"`
	RazonSocial      string `json:"razonSocial"      binding:"required"`
	CedulaJuridica   string `json:"cedulaJuridica"   binding:"required"`
	Telefono         string `json:"telefono"`
	Email            string `json:"email"            binding:"required"`
}

func (server *Server) createEmpresaCliente(ctx *gin.Context) {
	var req createEmpresaClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateEmpresaClienteParams{
		Razonsocial:    req.RazonSocial,
		Cedulajuridica: req.CedulaJuridica,
		Telefono:       toNullString(req.Telefono),
		Email:          req.Email,
	}

	result, err := server.dbtx.CreateEmpresaCliente(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idEmpresaCliente": lastId,
		"message":          "Empresa cliente creada",
	})
}

func (server *Server) getAllEmpresaCliente(ctx *gin.Context) {
	empresas, err := server.dbtx.GetAllEmpresaCliente(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, empresas)
}

func (server *Server) getEmpresaClienteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	empresa, err := server.dbtx.GetEmpresaClienteById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Empresa cliente no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, empresa)
}

func (server *Server) getEmpresaClienteByCedula(ctx *gin.Context) {
	cedula := ctx.Param("cedula")

	empresa, err := server.dbtx.GetEmpresaClienteByCedula(ctx, cedula)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Empresa cliente no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, empresa)
}

func (server *Server) updateEmpresaCliente(ctx *gin.Context) {
	var req updateEmpresaClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateEmpresaClienteParams{
		Razonsocial:      req.RazonSocial,
		Cedulajuridica:   req.CedulaJuridica,
		Telefono:         toNullString(req.Telefono),
		Email:            req.Email,
		Idempresacliente: req.IdEmpresaCliente,
	}

	_, err := server.dbtx.UpdateEmpresaCliente(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Empresa cliente actualizada"})
}

func (server *Server) deleteEmpresaCliente(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteEmpresaCliente(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Empresa cliente eliminada"})
}