package api;
import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)
type createEmailClienteRequest struct {
	Email        string    `json:"email"         binding:"required"`
	IdCliente    int32   `json:"idCliente"     binding:"required"`
}

type updateEmailClienteRequest struct {
	IdEmailCliente int32     `json:"idEmailCliente" binding:"required"`
	Email          string    `json:"email"          binding:"required"`
	IdCliente      int32    `json:"idCliente"      binding:"required"`
}

func (server *Server) createEmailCliente(ctx *gin.Context) {
	var req createEmailClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateEmailClienteParams{
		Email:        req.Email,
		Idcliente:    req.IdCliente,
	}
	emailCliente, err := server.dbtx.CreateEmailCliente(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := emailCliente.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllEmailCliente(ctx *gin.Context) {
	emailCliente, err := server.dbtx.GetAllEmailCliente(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, emailCliente)
}

func (server *Server) getEmailClienteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	emailCliente, err := server.dbtx.GetEmailClienteById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Email del cliente no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, emailCliente)
}

func (server *Server) updateEmailCliente(ctx *gin.Context) {
	var req updateEmailClienteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateEmailClienteParams{
		Idemailcliente: req.IdEmailCliente,
		Email:          req.Email,
		Idcliente:      req.IdCliente,
	}

	_, err := server.dbtx.UpdateEmailCliente(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Email del cliente actualizado correctamente"})
}

func (server *Server) deleteEmailCliente(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteEmailCliente(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Email del cliente eliminado"})
}
