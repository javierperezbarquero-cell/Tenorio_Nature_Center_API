package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createParticipanteRequest struct {
	IdReserva    int32     `json:"idReserva" binding:"required"`
	Nombre       string    `json:"nombre" binding:"required"`
	FechaNac     time.Time `json:"fechaNac" binding:"required"`
	Nacionalidad string    `json:"nacionalidad" binding:"required"`
	Telefono     string     `json:"telefono" binding:"required"`
}

type updateParticipanteRequest struct {
	IdParticipante int32     `json:"idParticipante" binding:"required"`
	IdReserva      int32     `json:"idReserva" binding:"required"`
	Nombre         string    `json:"nombre" binding:"required"`
	FechaNac       time.Time `json:"fechaNac" binding:"required"`
	Nacionalidad   string    `json:"nacionalidad" binding:"required"`
	Telefono       string     `json:"telefono" binding:"required"`
}

func (server *Server) createParticipante(ctx *gin.Context) {
	var req createParticipanteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateParticipanteParams{
		Idreserva:    req.IdReserva,
		Nombre:       req.Nombre,
		Fechanac:     req.FechaNac,
		Nacionalidad: req.Nacionalidad,
		Telefono:     req.Telefono,
	}

	result, err := server.dbtx.CreateParticipante(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idParticipante": lastId,
		"message":        "Participante creado",
	})
}

func (server *Server) getAllParticipantes(ctx *gin.Context) {
	participantes, err := server.dbtx.GetAllParticipante(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, participantes)
}

func (server *Server) getParticipanteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	participante, err := server.dbtx.GetParticipanteById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Participante no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, participante)
}

func (server *Server) updateParticipante(ctx *gin.Context) {
	var req updateParticipanteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateParticipanteParams{
		Idreserva:      req.IdReserva,
		Nombre:         req.Nombre,
		Fechanac:       req.FechaNac,
		Nacionalidad:   req.Nacionalidad,
		Telefono:       req.Telefono,
		Idparticipante: req.IdParticipante,
	}

	_, err := server.dbtx.UpdateParticipante(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Participante actualizado"})
}

func (server *Server) deleteParticipante(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteParticipante(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Participante eliminado"})
}
