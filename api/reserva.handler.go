package api
import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createReservaRequest struct {
	IdEstadoReserva int32 `json:"idEstadoReserva" binding:"required"`
}

type updateReservaRequest struct {
	IdReserva       int32 `json:"idReserva"       binding:"required"`
	IdEstadoReserva int32 `json:"idEstadoReserva" binding:"required"`
}

func (server *Server) createReserva(ctx *gin.Context) {
	var req createReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.dbtx.CreateReserva(ctx, req.IdEstadoReserva)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idReserva": lastId,
		"message":   "Reserva creada",
	})
}

func (server *Server) getAllReservas(ctx *gin.Context) {
	reservas, err := server.dbtx.GetAllReserva(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reservas)
}

func (server *Server) getReservaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	reserva, err := server.dbtx.GetReservaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Reserva no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reserva)
}

func (server *Server) updateReserva(ctx *gin.Context) {
	var req updateReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateReservaParams{
		Idestadoreserva: req.IdEstadoReserva,
		Idreserva:       req.IdReserva,
	}

	_, err := server.dbtx.UpdateReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva actualizada"})
}

func (server *Server) deleteReserva(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteReserva(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva eliminada"})
}