package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createIdiomaRequest struct {
	Nombre string `json:"nombre" 	binding:"required"`
}

type updateIdiomaRequest struct {
	IdIdioma int32  `json:"idIdioma" binding:"required"`
	Nombre   string `json:"nombre"   binding:"required"`
}

// POST: api/v1/idioma
func (server *Server) createIdioma(ctx *gin.Context) {
	var req createIdiomaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	idioma, err := server.dbtx.CreateIdioma(ctx, req.Nombre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := idioma.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

// GET: api/v1/idioma
func (server *Server) getAllIdioma(ctx *gin.Context) {
	idioma, err := server.dbtx.GetAllIdioma(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, idioma)
}

func (server *Server) getIdiomaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	idioma, err := server.dbtx.GetIdiomaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Idioma no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, idioma)
}

// PuT: api/v1/idioma
func (server *Server) updateIdioma(ctx *gin.Context) {
	var req updateIdiomaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateIdiomaParams{
		Nombre:   req.Nombre,
		Ididioma: req.IdIdioma,
	}

	_, err := server.dbtx.UpdateIdioma(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Idioma actualizado correctamente"})
}

// DELETE: api/v1/idioma
func (server *Server) deleteIdioma(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteIdioma(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Idioma eliminado"})
}
