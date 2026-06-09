package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createVehiculoRequest struct {
	IdChofer  int32  `json:"idChofer"  binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
	Capacidad int32  `json:"capacidad" binding:"required"`
	Modelo    string `json:"modelo"    binding:"required"`
}

type updateVehiculoRequest struct {
	IdVehiculo int32  `json:"idVehiculo" binding:"required"`
	IdChofer   int32  `json:"idChofer"   binding:"required"`
	Matricula  string `json:"matricula"  binding:"required"`
	Capacidad  int32  `json:"capacidad"  binding:"required"`
	Modelo     string `json:"modelo"     binding:"required"`
}

func (server *Server) createVehiculo(ctx *gin.Context) {
	var req createVehiculoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateVehiculoParams{
		Idchofer:  req.IdChofer,
		Matricula: req.Matricula,
		Capacidad: req.Capacidad,
		Modelo:    req.Modelo,
	}

	result, err := server.dbtx.CreateVehiculo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"idVehiculo": lastId, "message": "Vehículo creado"})
}

func (server *Server) getAllVehiculos(ctx *gin.Context) {
	vehiculos, err := server.dbtx.GetAllVehiculos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vehiculos)
}

func (server *Server) getVehiculoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	vehiculo, err := server.dbtx.GetVehiculoById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Vehículo no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vehiculo)
}

func (server *Server) updateVehiculo(ctx *gin.Context) {
	var req updateVehiculoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateVehiculoParams{
		Idchofer:   req.IdChofer,
		Matricula:  req.Matricula,
		Capacidad:  req.Capacidad,
		Modelo:     req.Modelo,
		Idvehiculo: req.IdVehiculo,
	}

	_, err := server.dbtx.UpdateVehiculo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Vehículo actualizado"})
}

func (server *Server) deleteVehiculo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteVehiculo(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Vehículo eliminado"})
}
