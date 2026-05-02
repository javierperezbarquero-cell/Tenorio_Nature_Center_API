package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createGuiaRequest struct {
	Nombre       string    `json:"nombre"		 	binding:"required"`
	FechaNac     time.Time `json:"fechanac" 	 	binding:"required"`
	Telefono     int32     `json:"telefono" 	 	binding:"required"`
	Nacionalidad string    `json:"nacionalidad" 	binding:"required"`
	Email        string    `json:"email" 		 	binding:"required"`
}

type updateGuiaRequest struct {
	IdGuia       int32     `json:"idGuia"	    binding:"required"`
	Nombre       string    `json:"nombre" 			binding:"required"`
	FechaNac     time.Time `json:"fechanac" 		binding:"required"`
	Telefono     int32     `json:"telefono" 		binding:"required"`
	Nacionalidad string    `json:"nacionalidad" 	binding:"required"`
	Email        string    `json:"email" 		 	binding:"required"`
}

// POST: api/v1/guia
func (server *Server) createGuia(ctx *gin.Context) {
	var req createGuiaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateGuiaParams{
		Nombre:       req.Nombre,
		Fechanac:     req.FechaNac,
		Telefono:     int32(req.Telefono),
		Nacionalidad: req.Nacionalidad,
		Email:        req.Email,
	}
	guia, err := server.dbtx.CreateGuia(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := guia.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

// GET: api/v1/guia
func (server *Server) getAllGuia(ctx *gin.Context) {
	guia, err := server.dbtx.GetAllGuia(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, guia)
}

func (server *Server) getGuiaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	guia, err := server.dbtx.GetGuiaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Guía no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, guia)
}

// PuT: api/v1/guia
func (server *Server) updateGuia(ctx *gin.Context) {
	var req updateGuiaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateGuiaParams{
		Nombre:       req.Nombre,
		Fechanac:     req.FechaNac,
		Telefono:     int32(req.Telefono),
		Nacionalidad: req.Nacionalidad,
		Email:        req.Email,
		Idguia:       req.IdGuia,
	}

	_, err := server.dbtx.UpdateGuia(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Guía actualizada correctamente"})
}

// DELETE: api/v1/guia
func (server *Server) deleteGuia(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteGuia(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Guía eliminada"})
}
