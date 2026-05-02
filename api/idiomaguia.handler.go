package api;
import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)
type createIdiomaGuiaRequest struct {
	IdGuia      int32   `json:"idGuia"       binding:"required"`
	IdIdioma    int32   `json:"idIdioma"     binding:"required"`
}

type updateIdiomaGuiaRequest struct {
	IdIdiomaGuia int32   `json:"idIdiomaGuia" binding:"required"`
	IdGuia       int32   `json:"idGuia"       binding:"required"`
	IdIdioma     int32   `json:"idIdioma"     binding:"required"`
}

func (server *Server) createIdiomaGuia(ctx *gin.Context) {
	var req createIdiomaGuiaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateIdiomaGuiaParams{
		Idguia:       req.IdGuia,
		Ididioma:     req.IdIdioma,
	}
	idiomaGuia, err := server.dbtx.CreateIdiomaGuia(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := idiomaGuia.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllIdiomaGuia(ctx *gin.Context) {
	idiomaGuia, err := server.dbtx.GetAllIdiomaGuia(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, idiomaGuia)
}

func (server *Server) getIdiomaGuiaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	idiomaGuia, err := server.dbtx.GetIdiomaGuiaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Idioma del guía no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, idiomaGuia)
}

func (server *Server) updateIdiomaGuia(ctx *gin.Context) {
	var req updateIdiomaGuiaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateIdiomaGuiaParams{
		Ididiomaguia:   req.IdIdiomaGuia,
		Idguia:         req.IdGuia,
		Ididioma:       req.IdIdioma,
	}

	_, err := server.dbtx.UpdateIdiomaGuia(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Idioma del guía actualizado correctamente"})
}

func (server *Server) deleteIdiomaGuia(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteIdiomaGuia(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Idioma del guía eliminado"})
}
