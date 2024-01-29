-- +goose Up
-- +goose StatementBegin
CREATE TABLE Invoicings (
    año int(4),
    mes int(2),
    tipo_comprobante varchar(4),
    secuencia_comprobante varchar(10),
    fecha_documento date,
    operacion varchar(50),
    puc_codigo_cuenta int(8),
    centro_costo varchar(3),
    punto_operacion int(1),
    tercero int(12),
    puc_nombre varchar(50),
    impuesto int,
    cheque varchar(255),
    debito float,
    credito float
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF ExISTS Invoicings
-- +goose StatementEnd
