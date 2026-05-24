package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createFacturaRequest struct {
	IdParticipante int32  `json:"idParticipante" binding:"required"`
	IdEstadoPago   int32  `json:"idEstadoPago"   binding:"required"`
	NumeroFactura  string `json:"numeroFactura"  binding:"required"`
	FechaFactura   string `json:"fechaFactura"   binding:"required"`
	MetodoPago     string `json:"metodoPago"     binding:"required"`
	Moneda         string `json:"moneda"         binding:"required"`
	FechaPago      string `json:"fechaPago"`
	Subtotal       string `json:"subtotal"       binding:"required"`
	Impuesto       string `json:"impuesto"       binding:"required"`
	Descuento      string `json:"descuento"      binding:"required"`
	PrecioTotal    string `json:"precioTotal"    binding:"required"`
}

type updateFacturaRequest struct {
	IdFactura      int32  `json:"idFactura"      binding:"required"`
	IdParticipante int32  `json:"idParticipante" binding:"required"`
	IdEstadoPago   int32  `json:"idEstadoPago"   binding:"required"`
	NumeroFactura  string `json:"numeroFactura"  binding:"required"`
	FechaFactura   string `json:"fechaFactura"   binding:"required"`
	MetodoPago     string `json:"metodoPago"     binding:"required"`
	Moneda         string `json:"moneda"         binding:"required"`
	FechaPago      string `json:"fechaPago"`
	Subtotal       string `json:"subtotal"       binding:"required"`
	Impuesto       string `json:"impuesto"       binding:"required"`
	Descuento      string `json:"descuento"      binding:"required"`
	PrecioTotal    string `json:"precioTotal"    binding:"required"`
}

func (server *Server) createFactura(ctx *gin.Context) {
	var req createFacturaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := validarCamposDecimalesFactura(req.Subtotal, req.Impuesto, req.Descuento, req.PrecioTotal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaFactura, err := parsearFecha(req.FechaFactura, "fechaFactura")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaPago, err := parsearFechaNullable(req.FechaPago, "fechaPago")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.CreateFacturaParams{
		Idparticipante: req.IdParticipante,
		Idestadopago:   req.IdEstadoPago,
		Numerofactura:  req.NumeroFactura,
		Fechafactura:   fechaFactura,
		Metodopago:     req.MetodoPago,
		Moneda:         req.Moneda,
		Fechapago:      fechaPago,
		Subtotal:       req.Subtotal,
		Impuesto:       req.Impuesto,
		Descuento:      req.Descuento,
		Preciototal:    req.PrecioTotal,
	}

	result, err := server.dbtx.CreateFactura(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idFactura": lastId,
		"message":   "Factura creada",
	})
}

func (server *Server) getAllFacturas(ctx *gin.Context) {
	facturas, err := server.dbtx.GetAllFactura(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, facturas)
}

func (server *Server) getFacturaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	factura, err := server.dbtx.GetFacturaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Factura no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, factura)
}

func (server *Server) updateFactura(ctx *gin.Context) {
	var req updateFacturaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := validarCamposDecimalesFactura(req.Subtotal, req.Impuesto, req.Descuento, req.PrecioTotal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaFactura, err := parsearFecha(req.FechaFactura, "fechaFactura")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fechaPago, err := parsearFechaNullable(req.FechaPago, "fechaPago")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.UpdateFacturaParams{
		Idparticipante: req.IdParticipante,
		Idestadopago:   req.IdEstadoPago,
		Numerofactura:  req.NumeroFactura,
		Fechafactura:   fechaFactura,
		Metodopago:     req.MetodoPago,
		Moneda:         req.Moneda,
		Fechapago:      fechaPago,
		Subtotal:       req.Subtotal,
		Impuesto:       req.Impuesto,
		Descuento:      req.Descuento,
		Preciototal:    req.PrecioTotal,
		Idfactura:      req.IdFactura,
	}

	_, err = server.dbtx.UpdateFactura(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Factura actualizada"})
}

func (server *Server) deleteFactura(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteFactura(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Factura eliminada"})
}
