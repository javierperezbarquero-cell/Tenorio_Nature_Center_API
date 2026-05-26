package api
 
import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"
 
	"github.com/gin-gonic/gin"
)
 
type createFacturaParticipanteRequest struct {
	IdFactura      int32 `json:"idFactura"      binding:"required"`
	IdParticipante int32 `json:"idParticipante" binding:"required"`
}
 
type updateFacturaParticipanteRequest struct {
	IdParticipante int32 `json:"idParticipante" binding:"required"`
	IdFactura      int32 `json:"idFactura"      binding:"required"`
}
 
func (server *Server) createFacturaParticipante(ctx *gin.Context) {
	var req createFacturaParticipanteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
 
	args := dto.CreateFacturaParticipanteParams{
		Idfactura:      req.IdFactura,
		Idparticipante: req.IdParticipante,
	}
 
	result, err := server.dbtx.CreateFacturaParticipante(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
 
	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idFacturaParticipante": lastId,
		"message":               "Participante asignado a la factura",
	})
}
 
func (server *Server) getParticipantesByFactura(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
 
	participantes, err := server.dbtx.GetParticipantesByFactura(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, participantes)
}
 
func (server *Server) getFacturaByParticipante(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
 
	facturaParticipante, err := server.dbtx.GetFacturaByParticipante(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Este participante no tiene factura asignada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facturaParticipante)
}
 
func (server *Server) updateFacturaParticipante(ctx *gin.Context) {
	var req updateFacturaParticipanteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
 
	args := dto.UpdateFacturaParticipanteParams{
		Idfactura:      req.IdFactura,
		Idparticipante: req.IdParticipante,
	}
 
	_, err := server.dbtx.UpdateFacturaParticipante(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Participante reasignado a nueva factura"})
}
 
func (server *Server) deleteFacturaParticipante(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
 
	_, err = server.dbtx.DeleteFacturaParticipante(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Asignación eliminada"})
}
 
func (server *Server) deleteFacturaParticipanteByFactura(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
 
	_, err = server.dbtx.DeleteFacturaParticipanteByFactura(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Participantes de la factura eliminados"})
}
 