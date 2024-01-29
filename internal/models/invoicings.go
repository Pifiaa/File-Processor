package models

// Representacion de las transacciónes
type Invoicings struct {
	Año                  string `json:"año"`
	Mes                  string `json:"mes"`
	TipoComprobante      string `json:"tipo_comprobante"`
	SecuenciaComprobante string `json:"secuencia_comprobante"`
	FechaDocumento       string `json:"fecha_documento"`
	Operacion            string `json:"operacion"`
	PucCodigoCuenta      string `json:"puc_codigo_cuenta"`
	CentroCosto          string `json:"centro_costo"`
	PuntoOperacion       string `json:"punto_operacion"`
	Tercero              string `json:"tercero"`
	PucNombre            string `json:"puc_nombre"`
	Impuesto             string `json:"impuesto"`
	Cheque               string `json:"cheque"`
	Debito               string `json:"debito"`
	Credito              string `json:"credito"`
}
