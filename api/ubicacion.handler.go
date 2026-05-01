package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createUbicacionRequest struct {
	Nombre    string `json:"nombre"        		binding:"required"`
	Direccion string `json:"direccion"     		binding:"required"`
}

type updateUbicacionRequest struct {
	IdUbicacion int32  `json:"idUbicacion"   binding:"required"`
	Nombre      string `json:"nombre"        binding:"required"`
	Direccion   string `json:"direccion"     binding:"required"`
}

func (server *Server) createUbicacion(ctx *gin.Context) {
	var req createUbicacionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateUbicacionParams{
		Nombre:    req.Nombre,
		Direccion: req.Direccion,
	}
	ubicacion, err := server.dbtx.CreateUbicacion(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := ubicacion.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllUbicacion(ctx *gin.Context) {
	ubicacion, err := server.dbtx.GetAllUbicacion(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ubicacion)
}

func (server *Server) getUbicacionById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ubicacion, err := server.dbtx.GetUbicacionById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Ubicación no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ubicacion)
}

func (server *Server) updateUbicacion(ctx *gin.Context) {
	var req updateUbicacionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateUbicacionParams{
		Nombre:      req.Nombre,
		Direccion:   req.Direccion,
		Idubicacion: req.IdUbicacion,
	}

	_, err := server.dbtx.UpdateUbicacion(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ubicación actualizada correctamente"})
}

func (server *Server) deleteUbicacion(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteUbicacion(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ubicación eliminada correctamente"})
}
