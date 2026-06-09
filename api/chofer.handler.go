package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createChoferRequest struct {
	Nombre        string `json:"nombre"        binding:"required"`
	Identificador string `json:"identificador" binding:"required"`
	FechaNac      string `json:"fechaNac"      binding:"required"`
	Telefono      string `json:"telefono"      binding:"required"`
	Email         string `json:"email"         binding:"required"`
	TipoLicencia  string `json:"tipoLicencia"  binding:"required"`
	Nacionalidad  string `json:"nacionalidad"  binding:"required"`
}

type updateChoferRequest struct {
	IdChofer      int32  `json:"idChofer"      binding:"required"`
	Nombre        string `json:"nombre"        binding:"required"`
	Identificador string `json:"identificador" binding:"required"`
	FechaNac      string `json:"fechaNac"      binding:"required"`
	Telefono      string `json:"telefono"      binding:"required"`
	Email         string `json:"email"         binding:"required"`
	TipoLicencia  string `json:"tipoLicencia"  binding:"required"`
	Nacionalidad  string `json:"nacionalidad"  binding:"required"`
}

func (server *Server) createChofer(ctx *gin.Context) {
	var req createChoferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fechaNac, err := parsearFecha(req.FechaNac, "fechaNac")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	args := dto.CreateChoferParams{
		Nombre:        req.Nombre,
		Identificador: req.Identificador,
		Fechanac:      fechaNac,
		Telefono:      req.Telefono,
		Email:         req.Email,
		Tipolicencia:  req.TipoLicencia,
		Nacionalidad:  req.Nacionalidad,
	}
	chofer, err := server.dbtx.CreateChofer(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := chofer.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllChofer(ctx *gin.Context) {
	chofer, err := server.dbtx.GetAllChofer(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, chofer)
}

func (server *Server) getChoferById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	chofer, err := server.dbtx.GetChoferById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Chofer no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, chofer)
}

func (server *Server) updateChofer(ctx *gin.Context) {
	var req updateChoferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fechaNac, err := parsearFecha(req.FechaNac, "fechaNac")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.UpdateChoferParams{
		Nombre:        req.Nombre,
		Identificador: req.Identificador,
		Fechanac:      fechaNac,
		Telefono:      req.Telefono,
		Email:         req.Email,
		Tipolicencia:  req.TipoLicencia,
		Nacionalidad:  req.Nacionalidad,
		Idchofer:      req.IdChofer,
	}

	_, err = server.dbtx.UpdateChofer(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Chofer actualizado correctamente"})
}

func (server *Server) deleteChofer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteChofer(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Chofer eliminado"})
}
