package domain

import "time"

type Evento struct {
	ID          int       `json:"id"`
	TipoSensor  string    `json:"tipo_sensor"` // Tipo de sensor (ej: "temperatura", "ultrasonido")
	Valor       float64   `json:"valor"`       // Valor registrado por el sensor
	Timestamp   time.Time `json:"timestamp"`   // Fecha y hora cuando se registró el evento
	CriadoEn    time.Time `json:"creado_en"`   // Fecha de creación
}
