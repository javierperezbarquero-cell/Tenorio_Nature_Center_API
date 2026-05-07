package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTourRequest struct {
	Nombre       string `json:"nombre"        binding:"required"`
	Descripcion  string `json:"descripcion"   binding:"required"`
	Horario      string `json:"horario"       binding:"required"`
	Duracion     int32  `json:"duracion"      binding:"required"`
	CuposMaximos int32  `json:"cuposMaximos"  binding:"required"`
	PrecioBase   string `json:"precioBase"    binding:"required"`
}

type updateTourRequest struct {
	IdTour       int32  `json:"idTour"        binding:"required"`
	Nombre       string `json:"nombre"        binding:"required"`
	Descripcion  string `json:"descripcion"   binding:"required"`
	Horario      string `json:"horario"       binding:"required"`
	Duracion     int32  `json:"duracion"      binding:"required"`
	CuposMaximos int32  `json:"cuposMaximos"  binding:"required"`
	PrecioBase   string `json:"precioBase"    binding:"required"`
}

func (server *Server) createTour(ctx *gin.Context) {
	var req createTourRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := validarCamposDecimalesTour(req.PrecioBase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.CreateTourParams{
		Nombre:       req.Nombre,
		Descripcion:  req.Descripcion,
		Horario:      req.Horario,
		Duracion:     req.Duracion,
		Cuposmaximos: req.CuposMaximos,
		Preciobase:   req.PrecioBase,
	}

	result, err := server.dbtx.CreateTour(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"idTour": id, "status": "Tour creado"})
}

func (server *Server) getAllTours(ctx *gin.Context) {
	tours, err := server.dbtx.GetAllTours(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tours)
}

func (server *Server) getTourById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	tour, err := server.dbtx.GetTourById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Tour no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tour)
}

func (server *Server) updateTour(ctx *gin.Context) {
	var req updateTourRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := validarCamposDecimalesTour(req.PrecioBase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.UpdateTourParams{
		Nombre:       req.Nombre,
		Descripcion:  req.Descripcion,
		Horario:      req.Horario,
		Duracion:     req.Duracion,
		Cuposmaximos: req.CuposMaximos,
		Preciobase:   req.PrecioBase,
		Idtour:       req.IdTour,
	}

	_, err := server.dbtx.UpdateTour(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Tour actualizado correctamente"})
}

func (server *Server) deleteTour(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteTour(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Tour eliminado exitosamente"})
}
