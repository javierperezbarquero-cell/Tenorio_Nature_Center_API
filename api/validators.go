package api

import (
    "fmt"
    "strconv"
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