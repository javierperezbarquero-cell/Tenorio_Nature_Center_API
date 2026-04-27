package api

import (
	"net/http"
	"rest/dto"
	"time"

	"github.com/gin-gonic/gin"
)

type createChoferRequest struct {
	Name         string    `json:"nombre"        binding:"required"`
	FechaNac     time.Time `json:"fechaNac"      binding:"required"`
	Telefono     int32     `json:"telefono"      binding:"required"`
	Email        string    `json:"email"         binding:"required"`
	TipoLicencia string    `json:"tipoLicencia"  binding:"required"`
	Nacionalidad string    `json:"nacionalidad"  binding:"required"`
}

func (server *Server) createChofer(ctx *gin.Context) {
	var req createChoferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateChoferParams{
		Nombre:       req.Name,
		Fechanac:     req.FechaNac,
		Telefono:     int32(req.Telefono),
		Email:        req.Email,
		Tipolicencia: req.TipoLicencia,
		Nacionalidad: req.Nacionalidad,
	}
	chofer, err := server.dbtx.CreateChofer(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := chofer.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAll(ctx *gin.Context) {
	chofer, err := server.dbtx.GetAllChofer(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, chofer)
}
