package application

import (
	"Eventos/src/Eventos/domain"
	"fmt"
	"time"
)

type ServicioEvento struct {
	repo domain.IEvento 
}

func NuevoServicioEvento(repo domain.IEvento) *ServicioEvento {
	return &ServicioEvento{repo}
}

func (s *ServicioEvento) CrearEvento(tipoSensor string, valor float64) (*domain.Evento, error) {
	evento := &domain.Evento{
		TipoSensor: tipoSensor,
		Valor:      valor,
		Timestamp:  time.Now(),
		CriadoEn:   time.Now(),
	}

	err := s.repo.Guardar(evento)
	if err != nil {
		return nil, fmt.Errorf("error al guardar evento: %v", err)
	}

	return evento, nil
}
