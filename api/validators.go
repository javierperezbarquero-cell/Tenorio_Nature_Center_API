package api

import (
    "fmt"
    "strconv"
	"time"
	"database/sql"
)

func validarDecimal(valor string, campo string) error {
    v, err := strconv.ParseFloat(valor, 64)
    if err != nil {
        return fmt.Errorf("%s debe ser un número válido (ej: 150.75)", campo)
    }
    if v < 0 {
        return fmt.Errorf("%s no puede ser negativo", campo)
    }
    return nil
}

func validarCamposDecimalesFactura(subtotal, impuesto, descuento, precioTotal string) error {
    campos := []struct {
        valor  string
        nombre string
    }{
        {subtotal, "subtotal"},
        {impuesto, "impuesto"},
        {descuento, "descuento"},
        {precioTotal, "precioTotal"},
    }

    for _, c := range campos {
        if err := validarDecimal(c.valor, c.nombre); err != nil {
            return err
        }
    }
    return nil
}

func validarCamposDecimalesTour(precioBase string) error {
    campos := []struct {
        valor  string
        nombre string
    }{
        {precioBase, "precioBase"},
    }

    for _, c := range campos {
        if err := validarDecimal(c.valor, c.nombre); err != nil {
            return err
        }
    }
    return nil
}

func parsearFecha(valor string, campo string) (time.Time, error) {
    fecha, err := time.Parse("2006-01-02", valor)
    if err != nil {
        return time.Time{}, fmt.Errorf("formato de %s inválido, use YYYY-MM-DD", campo)
    }
    return fecha, nil
}

func parsearFechaNullable(valor string, campo string) (sql.NullTime, error) {
    if valor == "" {
        return sql.NullTime{Valid: false}, nil
    }
    fecha, err := time.Parse("2006-01-02", valor)
    if err != nil {
        return sql.NullTime{}, fmt.Errorf("formato de %s inválido, use YYYY-MM-DD", campo)
    }
    return sql.NullTime{Time: fecha, Valid: true}, nil
}