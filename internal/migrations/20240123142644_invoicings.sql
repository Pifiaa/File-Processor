-- +goose Up
-- +goose StatementBegin
CREATE TABLE Invoicings (
    año int(4),
    mes int(2),
    tipo_comprobante varchar(255),
    secuencia_comprobante varchar(255),
    fecha_documento varchar(255),
    operacion varchar(255),
    puc_codigo_cuenta varchar(255),
    centro_costo varchar(255),
    punto_operacion varchar(255),
    tercero varchar(255),
    puc_nombre varchar(255),
    impuesto varchar(255),
    cheque varchar(255),
    debito VARCHAR(255),
    credito VARCHAR(255)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF ExISTS Invoicings
-- +goose StatementEnd
